package state

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mini-iac/internal/provider"
	"os"
)

type State struct {
	State map[string]provider.ResourceState
}

func (s *State) LoadState() error {
	fi, err := os.Open("./state.json")

	if os.IsNotExist(err) {
		_, err = os.Create("./state.json")
		if err != nil {
			return errors.New("Could not create file state.json")
		} else {
			s.State = make(map[string]provider.ResourceState)
			return nil
		}
	} else if err != nil {
		return errors.New("Could not open state.json Fatal error")
	}
	defer fi.Close()
	content, err := io.ReadAll(fi)
	if err != nil {
		return errors.New("Could not read state.json Fatal error")
	}
	err = json.Unmarshal(content, &s.State)
	if err != nil {
		return errors.New("Could not unmarshal json")
	}
	return nil

}

func (s *State) SaveState() error {
	//marshal provider.ResourceState into json, key with id, and write to state,json
	fi, err := os.Create("./state.json")
	if err != nil {
		return errors.New("Fatal error, could not open state.json")
	}
	defer fi.Close()
	js, err := json.MarshalIndent(s.State, "", "    ")
	if err != nil {
		return errors.New("Fatal error, could not marshal s.State")
	}

	_, err = fi.WriteString(string(js))
	if err != nil {
		return errors.New("Fatal error, could not write to state.json")
	}
	return nil

}

func (s *State) GetSingleState(id string) (*provider.ResourceState, error) {
	//take an id, search through map for value, if found return the resourceState, other wise return error.
	state, ok := s.State[id]
	if !ok {
		return nil, fmt.Errorf("Error, id %s does not exist in state", id)
	}

	return &state, nil
}

func (s *State) GetAllStates() ([]provider.ResourceState, error) {

	states := []provider.ResourceState{}

	for _, state := range s.State {
		states = append(states, state)
	}

	return states, nil
}
