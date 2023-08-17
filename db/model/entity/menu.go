// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Menu is the golang structure for table menu.
type Menu struct {
	Id         int         `json:"id"         description:""`
	Pid        int         `json:"pid"        description:""`
	Icon       string      `json:"icon"       description:""`
	BgImg      string      `json:"bgImg"      description:""`
	Name       string      `json:"name"       description:""`
	Path       string      `json:"path"       description:""`
	Sort       float64     `json:"sort"       description:""`
	Type       int         `json:"type"       description:"1normal 2menu 3 button"`
	Desc       string      `json:"desc"       description:""`
	FilePath   string      `json:"filePath"   description:""`
	Status     int         `json:"status"     description:""`
	CreatedAt  *gtime.Time `json:"createdAt"  description:""`
	UpdatedAt  *gtime.Time `json:"updatedAt"  description:""`
	Permission string      `json:"permission" description:"权限标识"`
}
