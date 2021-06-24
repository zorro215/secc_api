package service

import (
	"encoding/json"
	"secc_api/models"
)

//绑定亲属信息
func bindUserInfo(relationType uint8, blindInfo models.HealthBlindInfo) {
	decodeString, _ := json.Marshal(blindInfo)
	CallMethod("HealthAi.bind", relationType, decodeString)
}
