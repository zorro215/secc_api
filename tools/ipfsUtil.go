package tools

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"io/ioutil"
	"os"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

func getUri() string {
	ipfsHost, _ := beego.AppConfig.String("ipfs.host")
	ipfsPort, _ := beego.AppConfig.String("ipfs.port")
	if ipfsPort == "" {
		ipfsPort = "5001"
	}
	return ipfsHost + ":" + ipfsPort
}

//ipfs 上传字符串
func AddString(str string) {
	sh := shell.NewShell(getUri())
	cid, err := sh.Add(strings.NewReader(str))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("added %s", cid)
}

//ipfs 上传文件
func AddFile(path string) {
	sh := shell.NewShell(getUri())
	cid, err := sh.AddDir(path)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("added %s", cid)
}

func Cat(examplesHash string) {
	sh := shell.NewShell(getUri())
	rc, err := sh.Cat(examplesHash)
	if err != nil {
		fmt.Println(err.Error())
	}
	body, err := ioutil.ReadAll(rc)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("cat:", string(body))
}
