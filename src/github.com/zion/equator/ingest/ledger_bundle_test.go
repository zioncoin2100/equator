package ingest

import (
	"testing"

	"github.com/zion/equator/test"
)

func TestLedgerBundleLoad(t *testing.T) {
	tt := test.Start(t).ScenarioWithoutEquator("base")
	defer tt.Finish()

	bundle := &LedgerBundle{Sequence: 2}
	err := bundle.Load(tt.CoreSession())

	if tt.Assert.NoError(err) {
		tt.Assert.Equal(uint32(2), bundle.Header.Sequence)
		tt.Assert.Len(bundle.Transactions, 3)
		tt.Assert.Len(bundle.TransactionFees, 3)
	}
}
