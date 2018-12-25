// Package db provides helpers to connect to test databases.  It has no
// internal dependencies on horizon and so should be able to be imported by
// any horizon package.
package db

import (
	"fmt"
	"log"
	"testing"

	"github.com/jmoiron/sqlx"
	// pq enables postgres support
	_ "github.com/lib/pq"
	db "github.com/spn/go/support/db/dbtest"
)

var (
	coreDB     *sqlx.DB
	coreUrl    *string
	horizonDB  *sqlx.DB
	horizonUrl *string
)

// Horizon returns a connection to the horizon test database
func Horizon(t *testing.T) *sqlx.DB {
	if horizonDB != nil {
		return horizonDB
	}
	postgres := db.Postgres(t)
	horizonUrl = &postgres.DSN
	horizonDB = postgres.Open()
	return horizonDB
}

// HorizonURL returns the database connection the url any test
// use when connecting to the history/horizon database
func HorizonURL() string {
	if horizonUrl == nil {
		log.Panic(fmt.Errorf("Horizon not initialized"))
	}
	return *horizonUrl
}

// SpnCore returns a connection to the spn core test database
func SpnCore(t *testing.T) *sqlx.DB {
	if coreDB != nil {
		return coreDB
	}
	postgres := db.Postgres(t)
	coreUrl = &postgres.DSN
	coreDB = postgres.Open()
	return coreDB
}

// SpnCoreURL returns the database connection the url any test
// use when connecting to the spn-core database
func SpnCoreURL() string {
	if coreUrl == nil {
		log.Panic(fmt.Errorf("SpnCore not initialized"))
	}
	return *coreUrl
}
