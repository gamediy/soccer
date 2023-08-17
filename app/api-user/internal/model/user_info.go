package model

import "github.com/gogf/gf/v2/i18n/gi18n"

type UserInfo struct {
	Uid        float64 `json:"uid"`
	UidInt64   int64   `json:"uidInt64"`
	Account    string  `json:"account"`
	ClientIP   string  `json:"clientIp"`
	Lang       string  `json:"lang"`
	Pid        float64 `json:"pid"`
	ParentPath string  `json:"parentPath"`
	I18n       *gi18n.Manager
}
