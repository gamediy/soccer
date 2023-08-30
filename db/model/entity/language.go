// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Language is the golang structure for table language.
type Language struct {
	Id        uint64      `json:"id"        description:""`
	Name      string      `json:"name"      description:""`
	Desc      string      `json:"desc"      description:""`
	En        string      `json:"en"        description:"英语"`
	Fr        string      `json:"fr"        description:"法语"`
	Ja        string      `json:"ja"        description:"日语"`
	Hi        string      `json:"hi"        description:"印度语"`
	Zh        string      `json:"zh"        description:"中文"`
	Ko        string      `json:"ko"        description:"汉语"`
	Pt        string      `json:"pt"        description:"葡萄牙语"`
	Es        string      `json:"es"        description:"西班牙语"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	Hu        string      `json:"hu"        description:"匈牙利语"`
	Ar        string      `json:"ar"        description:"阿拉伯语"`
}
