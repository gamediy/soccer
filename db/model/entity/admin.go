// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Admin is the golang structure for table admin.
type Admin struct {
	Id        int         `json:"id"        description:""`
	Rid       int         `json:"rid"       description:"角色ID"`
	Uname     string      `json:"uname"     description:""`
	Pwd       string      `json:"pwd"       description:""`
	Nickname  string      `json:"nickname"  description:""`
	Email     string      `json:"email"     description:""`
	Phone     string      `json:"phone"     description:""`
	Status    int         `json:"status"    description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	KeySecret string      `json:"keySecret" description:""`
}
