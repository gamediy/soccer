// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Bank is the golang structure for table bank.
type Bank struct {
	Id       int    `json:"id"       description:""`
	Icon     string `json:"icon"     description:""`
	Currency string `json:"currency" description:""`
	Net      string `json:"net"      description:""`
	Protocol string `json:"protocol" description:""`
	Summary  string `json:"summary"  description:"简介"`
	Url      string `json:"url"      description:"官网"`
	Status   int    `json:"status"   description:""`
	Class    string `json:"class"    description:""`
}
