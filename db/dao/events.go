// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"star_net/db/dao/internal"
)

// internalEventsDao is internal type for wrapping internal DAO implements.
type internalEventsDao = *internal.EventsDao

// eventsDao is the data access object for table p_events.
// You can define custom methods on it to extend its functionality as you wish.
type eventsDao struct {
	internalEventsDao
}

var (
	// Events is globally public accessible object for table p_events operations.
	Events = eventsDao{
		internal.NewEventsDao(),
	}
)

// Fill with you ideas below.
