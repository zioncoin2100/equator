// Package db provides helpers to connect to test databases.  It has no
// internal dependencies on equator and so should be able to be imported by
// any equator package.
package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	// pq enables postgres support
	_ "github.com/lib/pq"
)

var (
	coreDB    *sqlx.DB
	equatorDB *sqlx.DB
)

const (
	// DefaultEquatorURL is the default postgres connection string for
	// equator's test database.
	DefaultEquatorURL = "postgres://localhost:5432/equator_test?sslmode=disable"

	// DefaultZioncoreURL is the default postgres connection string
	// for equator's test zion core database.
	DefaultZioncoreURL = "postgres://localhost:5432/zion-core_test?sslmode=disable"
)

// Equator returns a connection to the equator test database
func Equator() *sqlx.DB {
	if equatorDB != nil {
		return equatorDB
	}
	equatorDB = OpenDatabase(EquatorURL())
	return equatorDB
}

// EquatorURL returns the database connection the url any test
// use when connecting to the history/equator database
func EquatorURL() string {
	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		databaseURL = DefaultEquatorURL
	}

	return databaseURL
}

// OpenDatabase opens a database, panicing if it cannot
func OpenDatabase(dsn string) *sqlx.DB {
	db, err := sqlx.Open("postgres", dsn)

	if err != nil {
		log.Panic(err)
	}

	return db
}

// Zioncore returns a connection to the zion core test database
func Zioncore() *sqlx.DB {
	if coreDB != nil {
		return coreDB
	}
	coreDB = OpenDatabase(ZioncoreURL())
	return coreDB
}

// ZioncoreURL returns the database connection the url any test
// use when connecting to the zion-core database
func ZioncoreURL() string {
	databaseURL := os.Getenv("ZION_CORE_DATABASE_URL")

	if databaseURL == "" {
		databaseURL = DefaultZioncoreURL
	}

	return databaseURL
}
