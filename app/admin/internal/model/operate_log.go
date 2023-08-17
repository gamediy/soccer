package model

import "time"

type OperateLog struct {
	Account     string    `json:"account"`
	RoleName    string    `json:"roleName"`
	Ip          string    `json:"ip"`
	Request     string    `json:"request"`
	Response    string    `json:"response"`
	Method      string    `json:"method"`
	Path        string    `json:"path"`
	CreatedAt   time.Time `json:"createdAt"`
	ElapsedTime int64
}
