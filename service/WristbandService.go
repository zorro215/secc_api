package service

import (
	"secc_api/models"
	"secc_api/tools"
)

func QueryWristband(deviceId string) []models.Wristband {
	engine := tools.QueryEngine()
	applies := make([]models.Wristband, 0)
	where := engine.In("data_time", tools.GetWeekStr()).OrderBy("data_time")
	if deviceId != "" {
		where.And("device_id = ?", deviceId)
	}
	_ = where.Find(&applies)
	return applies
}

func AddWristband(w models.Wristband) {
	engine := tools.QueryEngine()
	_, _ = engine.InsertOne(w)
}
