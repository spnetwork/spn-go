package test

import (
	"github.com/spn/go/services/horizon/internal/test/scenarios"
)

func loadScenario(scenarioName string, includeHorizon bool) {
	spnCorePath := scenarioName + "-core.sql"
	horizonPath := scenarioName + "-horizon.sql"

	if !includeHorizon {
		horizonPath = "blank-horizon.sql"
	}

	scenarios.Load(SpnCoreDatabaseURL(), spnCorePath)
	scenarios.Load(DatabaseURL(), horizonPath)
}
