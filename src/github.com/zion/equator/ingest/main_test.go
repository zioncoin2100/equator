package ingest

import (
	"testing"

	"github.com/zion/go/network"
	"github.com/zion/equator/test"
)

func TestIngest(t *testing.T) {
	tt := test.Start(t).ScenarioWithoutEquator("kahuna")
	defer tt.Finish()

	s := ingest(tt)
	tt.Require.NoError(s.Err)
	tt.Assert.Equal(57, s.Ingested)

	// Test that re-importing fails
	s.Err = nil
	s.Run()
	tt.Require.Error(s.Err, "Reimport didn't fail as expected")

	// Test that re-importing fails with allowing clear succeeds
	s.Err = nil
	s.ClearExisting = true
	s.Run()
	tt.Require.NoError(s.Err, "Couldn't re-import, even with clear allowed")
}

func TestTick(t *testing.T) {
	tt := test.Start(t).ScenarioWithoutEquator("base")
	defer tt.Finish()
	sys := sys(tt)

	// ingest by tick
	s := sys.Tick()
	tt.Require.NoError(s.Err)
	tt.Require.Nil(sys.current)

	tt.UpdateLedgerState()
	s = sys.Tick()
	tt.Require.NotNil(s)
	tt.Require.NoError(s.Err)
}

func ingest(tt *test.T) *Session {
	sys := sys(tt)
	return sys.Tick()
}

func sys(tt *test.T) *System {
	return New(
		network.TestNetworkPassphrase,
		"",
		tt.CoreSession(),
		tt.EquatorSession(),
	)
}
