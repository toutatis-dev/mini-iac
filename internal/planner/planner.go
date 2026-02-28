package planner

import (
	"mini-iac/internal/ast"
	"mini-iac/internal/provider"
	"mini-iac/internal/state"
)

type Action string

const (
	CREATE Action = "create"
	UPDATE Action = "update"
	DELETE Action = "delete"
	NOOP   Action = "no-op"
)

type PlanItem struct {
	Action   Action
	Resource ast.Resource
}

type Plan struct {
	Items []PlanItem
}

func Planner(state *state.State, resource *ast.Manifest, provider provider.ResourceProvider) (Plan, error) {

}
