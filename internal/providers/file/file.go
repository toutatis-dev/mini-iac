package file

import (
	"errors"
	"mini-iac/internal/ast"
	"mini-iac/internal/provider"
	"os"
	"time"
)

//methods to implement
//Read(id string) (*ResourceState, error)
//Create(desired *ast.Resource) (*ResourceState, error)
//Update(id string, desired *ast.Resource) (*ResourceState, error)
//Delete(id string) error

type FileProvider struct {
}

func (f *FileProvider) Read(id string) (*provider.ResourceState, error) {

}

func (f *FileProvider) Create(desired *ast.Resource) (*provider.ResourceState, error) {
	//take desired state, create file, return up-to-date state
	fileName := desired.ResourceName
	fi, err := os.Create(fileName)
	if err != nil {
		return nil, errors.New("Could not create new file")
	}
	defer fi.Close()

	fileContent := desired.Properties["content"]
	_, err = fi.WriteString(fileContent)
	if err != nil {
		return nil, errors.New("Could not write to file")
	}

	state := provider.ResourceState{
		Provider:     desired.Provider,
		ResourceName: desired.ResourceName,
		Properties:   desired.Properties,
		Timestamp:    time.Now(),
	}

	return &state, nil

}

func (f *FileProvider) Update(id string, desired *ast.Resource) (*provider.ResourceState, error) {

}

func (f *FileProvider) Delete(id string) error {

}
