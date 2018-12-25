package horizon

import (
	"github.com/spn/go/services/horizon/internal/db2/core"
	"github.com/spn/go/services/horizon/internal/db2/history"
	"github.com/spn/go/support/db"
	"github.com/spn/go/support/log"
)

func initHorizonDb(app *App) {
	session, err := db.Open("postgres", app.config.DatabaseURL)

	if err != nil {
		log.Panic(err)
	}

	// Make sure MaxIdleConns is equal MaxOpenConns. In case of high variance
	// in number of requests closing and opening connections may slow down Horizon.
	session.DB.SetMaxIdleConns(app.config.MaxDBConnections)
	session.DB.SetMaxOpenConns(app.config.MaxDBConnections)

	app.historyQ = &history.Q{session}
}

func initCoreDb(app *App) {
	session, err := db.Open("postgres", app.config.SpnCoreDatabaseURL)

	if err != nil {
		log.Panic(err)
	}

	// Make sure MaxIdleConns is equal MaxOpenConns. In case of high variance
	// in number of requests closing and opening connections may slow down Horizon.
	session.DB.SetMaxIdleConns(app.config.MaxDBConnections)
	session.DB.SetMaxOpenConns(app.config.MaxDBConnections)
	app.coreQ = &core.Q{session}
}

func init() {
	appInit.Add("horizon-db", initHorizonDb, "app-context", "log")
	appInit.Add("core-db", initCoreDb, "app-context", "log")
}
