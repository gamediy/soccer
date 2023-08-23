// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Bank is the golang structure for table bank.
type Bank struct {
	Id       int    `json:"id"       description:""`
	Icon     string `json:"icon"     description:"图标"`
	Currency string `json:"currency" description:""`
	Net      string `json:"net"      description:"类型"`
	Protocol string `json:"protocol" description:"名称"`
	Summary  string `json:"summary"  description:"简介"`
	Url      string `json:"url"      description:"官网"`
	Status   int    `json:"status"   description:"状态 1开启 2关闭"`
	Class    string `json:"class"    description:""`
}
