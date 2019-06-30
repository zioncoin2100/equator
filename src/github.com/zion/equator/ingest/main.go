// Package ingest contains the ingestion system for equator.  This system takes
// data produced by the connected zion-core database, transforms it and
// inserts it into the equator database.
package ingest

import (
	"sync"

	sq "github.com/Masterminds/squirrel"
	metrics "github.com/rcrowley/go-metrics"
	"github.com/zion/go/support/db"
	"github.com/zion/equator/db2/core"
)

const (
	// CurrentVersion reflects the latest version of the ingestion
	// algorithm. As rows are ingested into the equator database, this version is
	// used to tag them.  In the future, any breaking changes introduced by a
	// developer should be accompanied by an increase in this value.
	//
	// Scripts, that have yet to be ported to this codebase can then be leveraged
	// to re-ingest old data with the new algorithm, providing a seamless
	// transition when the ingested data's structure changes.
	CurrentVersion = 9
)

// Cursor iterates through a zion core database's ledgers
type Cursor struct {
	// FirstLedger is the beginning of the range of ledgers (inclusive) that will
	// attempt to be ingested in this session.
	FirstLedger int32
	// LastLedger is the end of the range of ledgers (inclusive) that will
	// attempt to be ingested in this session.
	LastLedger int32
	// DB is the zion-core db that data is ingested from.
	DB *db.Session

	Metrics *IngesterMetrics

	// Err is the error that caused this iteration to fail, if any.
	Err error

	lg   int32
	tx   int
	op   int
	data *LedgerBundle
}

// EffectIngestion is a helper struct to smooth the ingestion of effects.  this
// struct will track what the correct operation to use and order to use when
// adding effects into an ingestion.
type EffectIngestion struct {
	Dest        *Ingestion
	OperationID int64
	err         error
	added       int
	parent      *Ingestion
}

// LedgerBundle represents a single ledger's worth of novelty created by one
// ledger close
type LedgerBundle struct {
	Sequence        int32
	Header          core.LedgerHeader
	TransactionFees []core.TransactionFee
	Transactions    []core.Transaction
}

// System represents the data ingestion subsystem of equator.
type System struct {
	// EquatorDB is the connection to the equator database that ingested data will
	// be written to.
	EquatorDB *db.Session

	// CoreDB is the zion-core db that data is ingested from.
	CoreDB *db.Session

	Metrics IngesterMetrics

	// Network is the passphrase for the network being imported
	Network string

	// ZionCoreURL is the http endpoint of the zion-core that data is being
	// ingested from.
	ZionCoreURL string

	// SkipCursorUpdate causes the ingestor to skip
	// reporting the "last imported ledger" cursor to
	// zion-core
	SkipCursorUpdate bool

	lock    sync.Mutex
	current *Session
}

// IngesterMetrics tracks all the metrics for the ingestion subsystem
type IngesterMetrics struct {
	ClearLedgerTimer  metrics.Timer
	IngestLedgerTimer metrics.Timer
	LoadLedgerTimer   metrics.Timer
}

// Ingestion receives write requests from a Session
type Ingestion struct {
	// DB is the sql connection to be used for writing any rows into the equator
	// database.
	DB *db.Session

	ledgers                  sq.InsertBuilder
	transactions             sq.InsertBuilder
	transaction_participants sq.InsertBuilder
	operations               sq.InsertBuilder
	operation_participants   sq.InsertBuilder
	effects                  sq.InsertBuilder
	accounts                 sq.InsertBuilder
	trades                   sq.InsertBuilder
}

// Session represents a single attempt at ingesting data into the history
// database.
type Session struct {
	Cursor    *Cursor
	Ingestion *Ingestion
	// Network is the passphrase for the network being imported
	Network string

	// ZionCoreURL is the http endpoint of the zion-core that data is being
	// ingested from.
	ZionCoreURL string

	// ClearExisting causes the session to clear existing data from the equator db
	// when the session is run.
	ClearExisting bool

	// SkipCursorUpdate causes the session to skip
	// reporting the "last imported ledger" cursor to
	// zion-core
	SkipCursorUpdate bool

	// Metrics is a reference to where the session should record its metric information
	Metrics *IngesterMetrics

	//
	// Results fields
	//

	// Err is the error that caused this session to fail, if any.
	Err error

	// Ingested is the number of ledgers that were successfully ingested during
	// this session.
	Ingested int
}

// New initializes the ingester, causing it to begin polling the zion-core
// database for now ledgers and ingesting data into the equator database.
func New(network string, coreURL string, core, equator *db.Session) *System {
	i := &System{
		Network:        network,
		ZionCoreURL: coreURL,
		EquatorDB:      equator,
		CoreDB:         core,
	}

	i.Metrics.ClearLedgerTimer = metrics.NewTimer()
	i.Metrics.IngestLedgerTimer = metrics.NewTimer()
	i.Metrics.LoadLedgerTimer = metrics.NewTimer()
	return i
}

// NewSession initialize a new ingestion session, from `first` to `last` using
// `i`.
func NewSession(first, last int32, i *System) *Session {
	hdb := i.EquatorDB.Clone()

	return &Session{
		Ingestion: &Ingestion{
			DB: hdb,
		},
		Cursor: &Cursor{
			FirstLedger: first,
			LastLedger:  last,
			DB:          i.CoreDB,
			Metrics:     &i.Metrics,
		},
		Network:          i.Network,
		ZionCoreURL:   i.ZionCoreURL,
		SkipCursorUpdate: i.SkipCursorUpdate,
		Metrics:          &i.Metrics,
	}
}
