package test

import (
	"github.com/zion/equator/test/scenarios"
)

func loadScenario(scenarioName string, includeEquator bool) {
	zionCorePath := scenarioName + "-core.sql"
	equatorPath := scenarioName + "-equator.sql"

	if !includeEquator {
		equatorPath = "blank-equator.sql"
	}

	scenarios.Load(ZionCoreDatabaseURL(), zionCorePath)
	scenarios.Load(DatabaseURL(), equatorPath)
}
