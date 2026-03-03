package planner

import (
	"errors"
	"mini-iac/internal/ast"
	"mini-iac/internal/provider"
	"mini-iac/internal/state"
	"reflect"
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

func Planner(state *state.State, resource *ast.Manifest, prov provider.ResourceProvider) (Plan, error) {

	plan := Plan{}
	for _, block := range resource.Blocks {

		res, ok := block.(*ast.Resource)
		if !ok {
			return Plan{}, errors.New("Could not cast block to *ast.Resource")

		}
		blockID := provider.ResourceID(res)

		resourceState, ok := state.State[blockID]
		if !ok {
			//add create to plan
			planItem := PlanItem{
				Action:   CREATE,
				Resource: *res,
			}
			plan.Items = append(plan.Items, planItem)

		} else {
			currentState, err := prov.Read(blockID)
			if err != nil {
				return Plan{}, errors.New(err.Error())
			}

			equal := reflect.DeepEqual(currentState.Properties, resourceState.Properties)
			if equal {
				//noop
				planItem := PlanItem{
					Action:   NOOP,
					Resource: *res,
				}
				plan.Items = append(plan.Items, planItem)
			} else {
				//update
				planItem := PlanItem{
					Action:   UPDATE,
					Resource: *res,
				}
				plan.Items = append(plan.Items, planItem)
			}
		}
	}
}
