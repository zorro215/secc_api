package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"pack.ag/amqp"
	"secc_api/models"
	_ "secc_api/routers"
	"secc_api/service"
	"time"
)

//参数说明，请参见AMQP客户端接入说明文档。
const accessKey = "LTAI5tKhBc8vP6yVRuPzD9na"
const accessSecret = "xLhbePbGTkTX8h13ZktMBLiY2AO7p2"
const consumerGroupId = "DEFAULT_GROUP"
const clientId = "123456"

//iotInstanceId：企业版实例请填写实例ID，公共实例请填空字符串""。
const iotInstanceId = ""

//接入域名，请参见AMQP客户端接入说明文档。
const host = "1563151931241081.iot-amqp.cn-shanghai.aliyuncs.com"

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	address := "amqps://" + host + ":5671"
	timestamp := time.Now().Nanosecond() / 1000000
	//userName组装方法，请参见AMQP客户端接入说明文档。
	userName := fmt.Sprintf("%s|authMode=aksign,signMethod=Hmacsha1,consumerGroupId=%s,authId=%s,iotInstanceId=%s,timestamp=%d|",
		clientId, consumerGroupId, accessKey, iotInstanceId, timestamp)
	stringToSign := fmt.Sprintf("authId=%s&timestamp=%d", accessKey, timestamp)
	hmacKey := hmac.New(sha1.New, []byte(accessSecret))
	hmacKey.Write([]byte(stringToSign))
	//计算签名，password组装方法，请参见AMQP客户端接入说明文档。
	password := base64.StdEncoding.EncodeToString(hmacKey.Sum(nil))

	amqpManager := &AmqpManager{
		address:  address,
		userName: userName,
		password: password,
	}

	//如果需要做接受消息通信或者取消操作，从Background衍生context。
	ctx := context.Background()

	amqpManager.startReceiveMessage(ctx)

	beego.Run()
}

//业务函数。用户自定义实现，该函数被异步执行，请考虑系统资源消耗情况。
func (am *AmqpManager) processMessage(message *amqp.Message) {
	fmt.Println("data received:", string(message.GetData()), " properties:", message.ApplicationProperties)

	// key：string类型，value：interface{}  类型能存任何数据类型
	var jsonObj map[string]interface{}
	_ = json.Unmarshal(message.GetData(), &jsonObj)
	params := jsonObj["items"].(map[string]interface{})
	dataId := params["DataId"].(map[string]interface{})
	dataTime := params["DataTime"].(map[string]interface{})
	deviceNo := params["DeviceNo"].(map[string]interface{})
	heartRate := params["HeartRate"].(map[string]interface{})

	w := models.Wristband{DataId: dataId["value"].(string), DeviceNo: deviceNo["value"].(string), DataTime: dataTime["value"].(string), HeartRate: uint16(heartRate["value"].(float64))}

	fmt.Println(w)
	service.AddWristband(w)
	wd := models.ConvertWristbandInfo(w)
	service.SaveWristbandInfo(wd)
}

type AmqpManager struct {
	address  string
	userName string
	password string
	client   *amqp.Client
	session  *amqp.Session
	receiver *amqp.Receiver
}

func (am *AmqpManager) startReceiveMessage(ctx context.Context) {

	childCtx, _ := context.WithCancel(ctx)
	err := am.generateReceiverWithRetry(childCtx)
	if nil != err {
		return
	}
	defer func() {
		_ = am.receiver.Close(childCtx)
		_ = am.session.Close(childCtx)
		_ = am.client.Close()
	}()

	for {
		//阻塞接受消息，如果ctx是background则不会被打断。
		message, err := am.receiver.Receive(ctx)

		if nil == err {
			go am.processMessage(message)
			_ = message.Accept()
		} else {
			fmt.Println("amqp receive data error:", err)

			//如果是主动取消，则退出程序。
			select {
			case <-childCtx.Done():
				return
			default:
			}

			//非主动取消，则重新建立连接。
			err := am.generateReceiverWithRetry(childCtx)
			if nil != err {
				return
			}
		}
	}
}

func (am *AmqpManager) generateReceiverWithRetry(ctx context.Context) error {
	//退避重连，从10ms依次x2，直到20s。
	duration := 10 * time.Millisecond
	maxDuration := 20000 * time.Millisecond
	times := 1

	//异常情况，退避重连。
	for {
		select {
		case <-ctx.Done():
			return amqp.ErrConnClosed
		default:
		}

		err := am.generateReceiver()
		if nil != err {
			time.Sleep(duration)
			if duration < maxDuration {
				duration *= 2
			}
			fmt.Println("amqp connect retry,times:", times, ",duration:", duration)
			times++
		} else {
			fmt.Println("amqp connect init success")
			return nil
		}
	}
}

//由于包不可见，无法判断Connection和Session状态，重启连接获取。
func (am *AmqpManager) generateReceiver() error {

	if am.session != nil {
		receiver, err := am.session.NewReceiver(
			amqp.LinkSourceAddress("/queue-name"),
			amqp.LinkCredit(20),
		)
		//如果断网等行为发生，Connection会关闭导致Session建立失败，未关闭连接则建立成功。
		if err == nil {
			am.receiver = receiver
			return nil
		}
	}

	//清理上一个连接。
	if am.client != nil {
		_ = am.client.Close()
	}

	client, err := amqp.Dial(am.address, amqp.ConnSASLPlain(am.userName, am.password))
	if err != nil {
		return err
	}
	am.client = client

	session, err := client.NewSession()
	if err != nil {
		return err
	}
	am.session = session

	receiver, err := am.session.NewReceiver(
		amqp.LinkSourceAddress("/queue-name"),
		amqp.LinkCredit(20),
	)
	if err != nil {
		return err
	}
	am.receiver = receiver

	return nil
}
