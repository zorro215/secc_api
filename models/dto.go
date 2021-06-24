package models

type WristbandDTO struct {
	Id        int64  `json:"id"`
	DataTime  string `json:"data_time"`
	DeviceId  string `json:"device_id"`
	HeartRate int64  `json:"heart_rate"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
