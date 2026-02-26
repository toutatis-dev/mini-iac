package provider

import (
	"fmt"
	"mini-iac/internal/ast"
	"time"
)

type ResourceProvider interface {
	Read(id string) (*ResourceState, error)
	Create(desired *ast.Resource) (*ResourceState, error)
	Update(id string, desired *ast.Resource) (*ResourceState, error)
	Delete(id string) error
}

type ResourceState struct {
	Provider     string
	ResourceName string
	Properties   map[string]string
	Timestamp    time.Time
}

func ResourceID(resource *ast.Resource) string {
	id := fmt.Sprintf("%s.%s", resource.Provider, resource.ResourceName)

	return id
}
