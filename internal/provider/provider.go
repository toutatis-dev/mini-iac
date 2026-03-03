package provider

import (
	"fmt"
	"mini-iac/internal/ast"
	"time"
)

type ResourceProvider interface {
	Read(desired *ast.Resource) (*ResourceState, error)
	Create(desired *ast.Resource) (*ResourceState, error)
	Update(desired *ast.Resource) (*ResourceState, error)
	Delete(resourceName string) error
}

type ResourceState struct {
	Provider     string            `json:"provider"`
	ResourceName string            `json:"resource_name"`
	Properties   map[string]string `json:"properties"`
	Timestamp    time.Time         `json:"timestamp"`
}

func ResourceID(resource *ast.Resource) string {
	id := fmt.Sprintf("%s.%s", resource.Provider, resource.ResourceName)

	return id
}
