package controllers

import (
	"encoding/json"
	"fmt"
	"secc_api/models"
	"secc_api/service"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about wristband
type WristbandController struct {
	beego.Controller
}

// @Title Create
// @Description create wristband
// @Param	body		body 	models.Wristband true		"The wristband content"
// @Success 200
// @Failure 403 body is empty
// @router / [post]
func (o *WristbandController) Post() {
	paramData := o.Ctx.Input.RequestBody
	w := models.Wristband{}
	err := json.Unmarshal(paramData, &w)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	//service.AddWristband(w)
}

// @Title Get
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
