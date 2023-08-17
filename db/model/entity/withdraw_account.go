// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// WithdrawAccount is the golang structure for table withdraw_account.
type WithdrawAccount struct {
	Id       int64  `json:"id"       description:""`
	Net      string `json:"net"      description:""`
	Protocol string `json:"protocol" description:""`
	Uid      int    `json:"uid"      description:""`
	Account  string `json:"account"  description:""`
	Address  string `json:"address"  description:""`
	Currency string `json:"currency" description:"currency"`
	Status   int    `json:"status"   description:""`
	Default  int    `json:"default"  description:"是否默认的"`
}
