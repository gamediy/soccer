package model

import "star_net/db/model/entity"

type AdminLoginLog struct {
	entity.AdminLoginLog
	Uname string `json:"uname"`
}
