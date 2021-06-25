package controllers

import (
	"encoding/json"
	"fmt"
	"secc_api/models"
	"secc_api/service"

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
// @Description find medical by deviceNo
// @Param	deviceNo		path 	string	true		"the deviceNo you want to get"
// @Success 200 {medical} models.Object
// @Failure 403 :deviceNo is empty
// @router /:deviceNo [get]
func (o *MedicalController) Get() {
	deviceNo := o.Ctx.Input.Param(":deviceNo")
	info := service.QueryMedical(deviceNo)
	o.Data["json"] = info
	_ = o.ServeJSON()
}
