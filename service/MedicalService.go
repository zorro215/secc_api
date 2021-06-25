package service

import (
	"secc_api/models"
	"secc_api/tools"
)

func QueryMedical(deviceNo string) []models.Medical {
	engine := tools.QueryEngine()
	applies := make([]models.Medical, 0)
	where := engine.In("data_time", tools.GetWeekStr()).OrderBy("data_time")
	if deviceNo != "" {
		where.And("device_no = ?", deviceNo)
	}
	_ = where.Find(&applies)
	return applies
}

func AddMedical(w models.Medical) {
	engine := tools.QueryEngine()
	_, _ = engine.InsertOne(w)
}
