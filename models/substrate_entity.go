package models

// HealthBlindInfo 绑定老人健康信息
type HealthBlindInfo struct {
	Name    string `json:"name"`
	IdCard  string `json:"id_card"`
	Height  uint32 `json:"height"`
	Weight  uint32 `json:"weight"`
	Chronic string `json:"chronic"`
}
