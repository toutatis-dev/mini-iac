package executor

import (
	"errors"
	"log"
	"mini-iac/internal/planner"
	"mini-iac/internal/provider"
	"mini-iac/internal/state"
)

func Execute(plan planner.Plan, prov provider.ResourceProvider, s *state.State) error {

	for _, plans := range plan.Items {
		switch plans.Action {
		case planner.CREATE:
			newState, err := prov.Create(plans.Resource)
			if err != nil {
				return errors.New(err.Error())
			}

			id := provider.ResourceID(plans.Resource)
			s.State[id] = *newState
		case planner.DELETE:
			resourceState := s.State[plans.ID]

			err := prov.Delete(resourceState.ResourceName)
			if err != nil {
				return errors.New(err.Error())
			}
			delete(s.State, plans.ID)
		case planner.NOOP:
			continue
		case planner.UPDATE:
			newState, err := prov.Update(plans.Resource)
			if err != nil {
				return errors.New(err.Error())
			}
			id := provider.ResourceID(plans.Resource)
			s.State[id] = *newState
		default:
			log.Println("Action not recognised")
		}
	}
	return nil
}
