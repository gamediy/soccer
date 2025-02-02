// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"star_net/db/dao/internal"
)

// internalSoccerOrderDao is internal type for wrapping internal DAO implements.
type internalSoccerOrderDao = *internal.SoccerOrderDao

// soccerOrderDao is the data access object for table o_soccer_order.
// You can define custom methods on it to extend its functionality as you wish.
type soccerOrderDao struct {
	internalSoccerOrderDao
}

var (
	// SoccerOrder is globally public accessible object for table o_soccer_order operations.
	SoccerOrder = soccerOrderDao{
		internal.NewSoccerOrderDao(),
	}
)

// Fill with you ideas below.
