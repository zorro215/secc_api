package models

import "time"

type SleepSign struct {
	DataId     string    `xorm:"pk comment('主键id') BIGINT"`
	DataTime   string    `xorm:"not null comment('数据时间') VARCHAR(32)"`
	DeviceNo   string    `xorm:"not null comment('设备标识') VARCHAR(128)"`
	HeartRate  uint8     `xorm:"comment('心率') BIGINT"`
	BreathRate uint8     `xorm:"comment('呼吸') BIGINT"`
	CreatedAt  time.Time `xorm:"created"`
	UpdatedAt  time.Time `xorm:"updated"`
	DeletedAt  time.Time `xorm:"deleted"`
}
