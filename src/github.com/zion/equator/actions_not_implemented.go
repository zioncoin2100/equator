package equator

import "github.com/zion/equator/render/problem"

// NotImplementedAction renders a NotImplemented prblem
type NotImplementedAction struct {
	Action
}

// JSON is a method for actions.JSON
func (action *NotImplementedAction) JSON() {
	problem.Render(action.Ctx, action.W, problem.NotImplemented)
}
