package ast

type Resource struct {
	Provider     string
	ResourceName string
	Properties   map[string]string
}

type Manifest struct {
	Blocks []Resource
}
