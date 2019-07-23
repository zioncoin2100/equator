package equator

import (
	"encoding/json"
	"testing"

	"github.com/zion/equator/resource"
	"github.com/zion/equator/test"
)

func TestRootAction(t *testing.T) {
	ht := StartHTTPTest(t, "base")
	defer ht.Finish()

	server := test.NewStaticMockServer(`{
			"info": {
				"network": "test",
				"build": "test-core",
				"protocol_version": 4
			}
		}`)
	defer server.Close()

	ht.App.equatorVersion = "test-equator"
	ht.App.config.ZioncoreURL = server.URL

	w := ht.Get("/")
	if ht.Assert.Equal(200, w.Code) {
		var actual resource.Root
		err := json.Unmarshal(w.Body.Bytes(), &actual)
		ht.Require.NoError(err)
		ht.Assert.Equal("test-equator", actual.EquatorVersion)
		ht.Assert.Equal("test-core", actual.ZioncoreVersion)
	}
}
