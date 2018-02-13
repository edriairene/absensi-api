package models

type Presence struct {
	Nim int `json:"nim"`
	Presence int `json:"presence"`
	PresTime int64 `json:"pres_time"`
}