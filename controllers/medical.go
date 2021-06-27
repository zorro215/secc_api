package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"secc_api/models"
	"secc_api/service"
	"secc_api/tools"

	beego "github.com/beego/beego/v2/server/web"
)

// MedicalController Operations about medical
type MedicalController struct {
	beego.Controller
}

// Post @Title Create
// @Description create medical
// @Param	body		body 	models.Medical true		"The medical content"
// @Success 200
// @Failure 403 body is empty
// @router / [post]
func (o *MedicalController) Post() {
	paramData := o.Ctx.Input.RequestBody
	var m models.Medical
	err1 := json.Unmarshal(paramData, &m)
	if err1 != nil {
		fmt.Println("json.Unmarshal is err1:", err1.Error())
	} else {
		service.AddMedical(m)
	}
	var md models.MedicalInfo
	err2 := json.Unmarshal(paramData, &md)
	if err2 != nil {
		fmt.Println("json.Unmarshal is err2:", err2.Error())
	} else {
		service.SaveMedicalInfo(md)
	}
}

// Get @Title Get
// @Description find medical by idCard
// @Param	idCard		path 	string	true		"the idCard you want to get"
// @Success 200 {medical} models.Medical
// @Failure 403 :idCard is empty
// @router /:idCard [get]
func (o *MedicalController) Get() {
	deviceNo := o.Ctx.Input.Param(":idCard")
	info := service.QueryMedical(deviceNo)
	o.Data["json"] = info
	_ = o.ServeJSON()
}

// UploadFile @Title UploadFile
// @Description UploadFile
// @Param	body		body 	multipart.File true		"The UploadFile content"
// @Success 200 {string} hash
// @Failure 403 body is empty
// @router /UploadFile [post]
func (o *MedicalController) UploadFile() {
	f, _, err := o.GetFile("uploadname")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, f)
	var hash string
	hash, _, _ = tools.UploadToIpfs(buf.Bytes())
	o.Data["string"] = hash
	_ = o.ServeJSON()
}
