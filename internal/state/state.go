package state

import (
	"encoding/json"
	"errors"
	"io"
	"mini-iac/internal/provider"
	"os"
)

type State struct {
	State map[string]provider.ResourceState
}

func (s *State) LoadState() error {
	fi, err := os.Open("./state.json")
	defer fi.Close()
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

}

func (s *State) GetSingleState(id string) (provider.ResourceState, error) {

}

func (s *State) GetAllStates() ([]provider.ResourceState, error) {

}
