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
// @Param	body		body 	models.Wristband	true		"The wristband content"
// @Success 200
// @Failure 403 body is empty
// @router / [post]
func (o *WristbandController) Post() {
	paramData := o.Ctx.Input.RequestBody
	var w models.Wristband
	err := json.Unmarshal(paramData, &w)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	service.AddWristband(w)
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

//// @Title GetAll
//// @Description get all wristbands
//// @Success 200 {wristband} models.Object
//// @Failure 403 :wristbandId is empty
//// @router / [get]
//func (o *WristbandController) GetAll() {
//	obs := models.GetAll()
//	o.Data["json"] = obs
//	o.ServeJSON()
//}
//
//// @Title Update
//// @Description update the wristband
//// @Param	wristbandId		path 	string	true		"The wristbandid you want to update"
//// @Param	body		body 	models.Object	true		"The body"
//// @Success 200 {wristband} models.Object
//// @Failure 403 :wristbandId is empty
//// @router /:wristbandId [put]
//func (o *WristbandController) Put() {
//	wristbandId := o.Ctx.Input.Param(":wristbandId")
//	var ob models.Object
//	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
//
//	err := models.Update(wristbandId, ob.Score)
//	if err != nil {
//		o.Data["json"] = err.Error()
//	} else {
//		o.Data["json"] = "update success!"
//	}
//	o.ServeJSON()
//}
//
//// @Title Delete
//// @Description delete the wristband
//// @Param	wristbandId		path 	string	true		"The wristbandId you want to delete"
//// @Success 200 {string} delete success!
//// @Failure 403 wristbandId is empty
//// @router /:wristbandId [delete]
//func (o *WristbandController) Delete() {
//	wristbandId := o.Ctx.Input.Param(":wristbandId")
//	models.Delete(wristbandId)
//	o.Data["json"] = "delete success!"
//	o.ServeJSON()
//}
