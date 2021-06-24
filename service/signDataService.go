package service

import (
	"encoding/json"
	"github.com/centrifuge/go-substrate-rpc-client/v3/types"
	"secc_api/models"
)

// SaveSleepSignInfo 保存睡眠体征数据/**
func SaveSleepSignInfo(info models.SleepSignInfo) {
	decodeString, _ := json.Marshal(info)
	CallMethod("SignData.saveSleepSignInfo", decodeString)
}

// SaveMedicalInfo 保存体检报告数据/**
func SaveMedicalInfo(info models.MedicalInfo) {
	CallMethod("SignData.saveMedicalInfo", info.FileHash, info.IdCard)
}

// SaveSleepReportInfo 保存睡眠报告数据/**
func SaveSleepReportInfo(info models.SleepReportInfo) {
	decodeString, _ := json.Marshal(info)
	CallMethod("SignData.saveSleepReportInfo", decodeString)
}

// SaveWristbandInfo 保存手环心率数据/**
func SaveWristbandInfo(info models.WristbandInfo) {
	decodeString, _ := json.Marshal(info)
	toString := types.HexEncodeToString(decodeString)
	CallMethod("SignData.saveWristbandInfo", toString)
}
