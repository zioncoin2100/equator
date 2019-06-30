package test

import (
	"io"

	"encoding/json"

	"github.com/zion/go/support/db"
	"github.com/zion/equator/ledger"
)

// CoreSession returns a db.Session instance pointing at the zion core test database
func (t *T) CoreSession() *db.Session {
	return &db.Session{
		DB:  t.CoreDB,
		Ctx: t.Ctx,
	}
}

// Finish finishes the test, logging any accumulated equator logs to the logs
// output
func (t *T) Finish() {
	RestoreLogger()
	// Reset cached ledger state
	ledger.SetState(ledger.State{})

	if t.LogBuffer.Len() > 0 {
		t.T.Log("\n" + t.LogBuffer.String())
	}
}

// EquatorSession returns a db.Session instance pointing at the equator test
// database
func (t *T) EquatorSession() *db.Session {
	return &db.Session{
		DB:  t.EquatorDB,
		Ctx: t.Ctx,
	}
}

// Scenario loads the named sql scenario into the database
func (t *T) Scenario(name string) *T {
	LoadScenario(name)
	t.UpdateLedgerState()
	return t
}

// ScenarioWithoutEquator loads the named sql scenario into the database
func (t *T) ScenarioWithoutEquator(name string) *T {
	LoadScenarioWithoutEquator(name)
	t.UpdateLedgerState()
	return t
}

// UnmarshalPage populates dest with the records contained in the json-encoded
// page in r.
func (t *T) UnmarshalPage(r io.Reader, dest interface{}) {
	var env struct {
		Embedded struct {
			Records json.RawMessage `json:"records"`
		} `json:"_embedded"`
	}

	err := json.NewDecoder(r).Decode(&env)
	t.Require.NoError(err, "failed to decode page")

	err = json.Unmarshal(env.Embedded.Records, dest)
	t.Require.NoError(err, "failed to decode records")
}

// UpdateLedgerState updates the cached ledger state (or panicing on failure).
func (t *T) UpdateLedgerState() {
	var next ledger.State

	err := t.CoreSession().GetRaw(&next, `
		SELECT
			COALESCE(MIN(ledgerseq), 0) as core_elder,
			COALESCE(MAX(ledgerseq), 0) as core_latest
		FROM ledgerheaders
	`)

	if err != nil {
		panic(err)
	}

	err = t.EquatorSession().GetRaw(&next, `
			SELECT
				COALESCE(MIN(sequence), 0) as history_elder,
				COALESCE(MAX(sequence), 0) as history_latest
			FROM history_ledgers
		`)

	if err != nil {
		panic(err)
	}

	ledger.SetState(next)
	return
}
