// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OperationLog is the golang structure for table operation_log.
type OperationLog struct {
	Id        int         `json:"id"        description:""`
	Uid       int         `json:"uid"       description:""`
	Content   string      `json:"content"   description:""`
	Response  string      `json:"response"  description:""`
	Method    string      `json:"method"    description:""`
	Uri       string      `json:"uri"       description:""`
	Ip        string      `json:"ip"        description:""`
	UseTime   int         `json:"useTime"   description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	MenuName  string      `json:"menuName"  description:""`
}
