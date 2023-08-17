package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"star_net/db/model/entity"
	"time"
)

type UserInfo struct {
	Uid       float64   `json:"uid"`
	AdminId   int64     `json:"adminId"`
	Account   string    `json:"account"`
	Email     string    `json:"email"`
	RuleName  string    `json:"ruleName"`
	RuleId    int       `json:"ruleId"`
	ClientIp  string    `json:"clientIp"`
	NIckName  string    `json:"nickName"`
	Phone     string    `json:"phone"`
	LoginTime time.Time `json:"loginTime"`
}
type UserLoginLog struct {
	Id        uint64      `json:"id"        description:""`
	Uid       uint64      `json:"uid"       description:""`
	Uname     string      `json:"uname"`
	Ip        string      `json:"ip"        description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
}
type User struct {
	User   *entity.User   `json:"user"`
	Wallet *entity.Wallet `json:"wallet"`
}
