package horizon

import (
	"log"

	"github.com/spn/go/services/horizon/internal/ingest"
)

func initIngester(app *App) {
	if !app.config.Ingest {
		return
	}

	if app.networkPassphrase == "" {
		log.Fatal("Cannot start ingestion without network passphrase.  Please confirm connectivity with spn-core.")
	}

	app.ingester = ingest.New(
		app.networkPassphrase,
		app.config.SpnCoreURL,
		app.CoreSession(nil),
		app.HorizonSession(nil),
		ingest.Config{
			EnableAssetStats: app.config.EnableAssetStats,
		},
	)

	app.ingester.SkipCursorUpdate = app.config.SkipCursorUpdate
	app.ingester.HistoryRetentionCount = app.config.HistoryRetentionCount
}

func init() {
	appInit.Add("ingester", initIngester, "app-context", "log", "horizon-db", "core-db", "spnCoreInfo")
}
