// Package ledger provides useful utilities concerning ledgers within zion,
// specifically a central location to store a cached snapshot of the state of
// both equator's and zion-core's view of the ledger.  This package is
// intended to be at the lowest levels of equator's dependency tree, please keep
// it free of dependencies to other equator packages.
package ledger

import (
	"sync"
)

// State represents a snapshot of both equator's and zion-core's view of the
// ledger.
type State struct {
	CoreLatest    int32 `db:"core_latest"`
	CoreElder     int32 `db:"core_elder"`
	HistoryLatest int32 `db:"history_latest"`
	HistoryElder  int32 `db:"history_elder"`
}

// CurrentState returns the cached snapshot of ledger state
func CurrentState() State {
	lock.RLock()
	ret := current
	lock.RUnlock()
	return ret
}

// SetState updates the cached snapshot of the ledger state
func SetState(next State) {
	lock.Lock()
	current = next
	lock.Unlock()
}

var current State
var lock sync.RWMutex
