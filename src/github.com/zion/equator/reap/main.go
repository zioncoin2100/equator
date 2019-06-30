// Package reap contains the history reaping subsystem for equator.  This system
// is designed to remove data from the history database such that it does not
// grow indefinitely.  The system can be configured with a number of ledgers to
// maintain at a minimum.
package reap

import (
	"time"

	"github.com/zion/go/support/db"
)

// System represents the history reaping subsystem of equator.
type System struct {
	EquatorDB      *db.Session
	RetentionCount uint

	nextRun time.Time
}

// New initializes the reaper, causing it to begin polling the zion-core
// database for now ledgers and ingesting data into the equator database.
func New(retention uint, equator *db.Session) *System {
	r := &System{
		EquatorDB:      equator,
		RetentionCount: retention,
	}

	r.nextRun = time.Now().Add(1 * time.Hour)
	return r
}
