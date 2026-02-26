package ast

type Block interface {
	isBlock()
}

type Resource struct {
	Provider     string
	ResourceName string
	Properties   map[string]string
}

type Manifest struct {
	Blocks []Block
}

func (r *Resource) isBlock() {}
