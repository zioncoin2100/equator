package equator

import (
	"github.com/PuerkitoBio/throttled"
	"github.com/sirupsen/logrus"
)

// Config is the configuration for equator.  It get's populated by the
// app's main function and is provided to NewApp.
type Config struct {
	DatabaseURL            string
	ZioncoreDatabaseURL string
	ZioncoreURL         string
	Port                   int
	RateLimit              throttled.Quota
	RedisURL               string
	LogLevel               logrus.Level
	SentryDSN              string
	LogglyHost             string
	LogglyToken            string
	FriendbotSecret        string
	// TLSCert is a path to a certificate file to use for equator's TLS config
	TLSCert string
	// TLSKey is the path to a private key file to use for equator's TLS config
	TLSKey string
	// Ingest is a boolean that indicates whether or not this equator instance
	// should run the data ingestion subsystem.
	Ingest bool
	// HistoryRetentionCount represents the minimum number of ledgers worth of
	// history data to retain in the equator database. For the purposes of
	// determining a "retention duration", each ledger roughly corresponds to 10
	// seconds of real time.
	HistoryRetentionCount uint

	// StaleThreshold represents the number of ledgers a history database may be
	// out-of-date by before equator begins to respond with an error to history
	// requests.
	StaleThreshold uint

	// SkipCursorUpdate causes the ingestor to skip reporting the "last imported
	// ledger" state to zion-core.
	SkipCursorUpdate bool
}
