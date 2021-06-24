package controllers

import (
	"encoding/json"
	"fmt"
	"secc_api/models"
	"secc_api/service"

	beego "github.com/beego/beego/v2/server/web"
)

// WristbandController Operations about wristband
type WristbandController struct {
	beego.Controller
}

// Post @Title Create
// @Description create wristband
// @Param	body		body 	models.Wristband true		"The wristband content"
// @Success 200
// @Failure 403 body is empty
// @router / [post]
func (o *WristbandController) Post() {
	paramData := o.Ctx.Input.RequestBody
	var w models.Wristband
	err := json.Unmarshal(paramData, &w)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	} else {
		service.AddWristband(w)
	}
}

// Get @Title Get
// @Description find wristband by deviceId
// @Param	deviceId		path 	string	true		"the deviceId you want to get"
// @Success 200 {wristband} models.Object
// @Failure 403 :deviceId is empty
// @router /:deviceId [get]
func (o *WristbandController) Get() {
	deviceId := o.Ctx.Input.Param(":deviceId")
	info := service.QueryWristband(deviceId)
	o.Data["json"] = info
	_ = o.ServeJSON()
}

// Bind @Title Bind
// @Description bind user info
// @Param	body		body 	models.HealthBindInfoDTO true		"The bindIndo content"
// @Success 200
// @Failure 403 body is empty
// @router /:bind [post]
func (o *WristbandController) Bind() {
	paramData := o.Ctx.Input.RequestBody
	w := models.HealthBindInfoDTO{}
	err := json.Unmarshal(paramData, &w)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	service.BindUserInfo(w)
}
