package equator

import (
	"net/http"

	"github.com/zion/equator/db2/core"
	"github.com/zion/equator/db2/history"
	"github.com/zion/equator/txsub"
	results "github.com/zion/equator/txsub/results/db"
	"github.com/zion/equator/txsub/sequence"
)

func initSubmissionSystem(app *App) {
	cq := &core.Q{Session: app.CoreSession(nil)}

	app.submitter = &txsub.System{
		Pending:         txsub.NewDefaultSubmissionList(),
		Submitter:       txsub.NewDefaultSubmitter(http.DefaultClient, app.config.ZionCoreURL),
		SubmissionQueue: sequence.NewManager(),
		Results: &results.DB{
			Core:    cq,
			History: &history.Q{Session: app.EquatorSession(nil)},
		},
		Sequences:         cq.SequenceProvider(),
		NetworkPassphrase: app.networkPassphrase,
	}
}

func init() {
	appInit.Add("txsub", initSubmissionSystem, "app-context", "log", "equator-db", "core-db")
}
