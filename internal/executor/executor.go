package executor

import (
	"mini-iac/internal/planner"
	"mini-iac/internal/provider"
	"mini-iac/internal/state"
)

func Execute(plan planner.Plan, prov provider.ResourceProvider, s *state.State) error {
}
