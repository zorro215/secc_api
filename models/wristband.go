package models

import "time"

type Wristband struct {
	Id        int64     `xorm:"pk autoincr comment('主键id') BIGINT"`
	DataTime  string    `xorm:"not null comment('数据时间') VARCHAR(50)"`
	DeviceId  string    `xorm:"not null comment('设备标识') VARCHAR(50)"`
	HeartRate int64     `xorm:"comment('心率') BIGINT"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}