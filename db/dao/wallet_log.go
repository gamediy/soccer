// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"star_net/db/dao/internal"
)

// internalWalletLogDao is internal type for wrapping internal DAO implements.
type internalWalletLogDao = *internal.WalletLogDao

// walletLogDao is the data access object for table u_wallet_log.
// You can define custom methods on it to extend its functionality as you wish.
type walletLogDao struct {
	internalWalletLogDao
}

var (
	// WalletLog is globally public accessible object for table u_wallet_log operations.
	WalletLog = walletLogDao{
		internal.NewWalletLogDao(),
	}
)

// Fill with you ideas below.
