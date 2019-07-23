package test

import (
	"github.com/zion/equator/test/scenarios"
)

func loadScenario(scenarioName string, includeEquator bool) {
	ZioncorePath := scenarioName + "-core.sql"
	equatorPath := scenarioName + "-equator.sql"

	if !includeEquator {
		equatorPath = "blank-equator.sql"
	}

	scenarios.Load(ZioncoreDatabaseURL(), ZioncorePath)
	scenarios.Load(DatabaseURL(), equatorPath)
}
