package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"star_net/consts"
	"star_net/db/model/entity"
)

type Menu struct {
	Id         int           `json:"id"         description:""`
	Pid        int           `json:"pid"        description:""`
	Icon       string        `json:"icon"       description:""`
	BgImg      string        `json:"bgImg"      description:""`
	Name       string        `json:"name"       description:""`
	Path       string        `json:"path"       description:""`
	Sort       float64       `json:"sort"       description:""`
	Type       int           `json:"type"       description:"1normal 2group"`
	Desc       string        `json:"desc"       description:""`
	FilePath   string        `json:"filePath"   description:""`
	Status     int           `json:"status"     description:""`
	CreatedAt  *gtime.Time   `json:"createdAt"  description:""`
	UpdatedAt  *gtime.Time   `json:"updatedAt"  description:""`
	Permission string        `json:"permission" description:"权限标识"`
	Api        []*entity.Api `json:"api"`
	Children   []*Menu       `json:"children"`
}

func (d *Menu) AddImgPrefix() {
	if d.Icon != "" {
		d.Icon = consts.ImgPrefix + d.Icon
	}
	if d.BgImg != "" {
		d.Icon = consts.ImgPrefix + d.Icon
	}
}
