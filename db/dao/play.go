// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"star_net/db/dao/internal"
)

// internalPlayDao is internal type for wrapping internal DAO implements.
type internalPlayDao = *internal.PlayDao

// playDao is the data access object for table p_play.
// You can define custom methods on it to extend its functionality as you wish.
type playDao struct {
	internalPlayDao
}

var (
	// Play is globally public accessible object for table p_play operations.
	Play = playDao{
		internal.NewPlayDao(),
	}
)

// Fill with you ideas below.
