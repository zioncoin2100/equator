package equator

import (
	"github.com/zion/equator/ledger"
	"github.com/zion/equator/render/hal"
	"github.com/zion/equator/resource"
)

// RootAction provides a summary of the equator instance and links to various
// useful endpoints
type RootAction struct {
	Action
}

// JSON renders the json response for RootAction
func (action *RootAction) JSON() {
	action.App.UpdateZioncoreInfo()

	var res resource.Root
	res.Populate(
		action.Ctx,
		ledger.CurrentState(),
		action.App.equatorVersion,
		action.App.coreVersion,
		action.App.networkPassphrase,
		action.App.protocolVersion,
	)

	hal.Render(action.W, res)
}
