package registry

import (
	"fmt"
)

type Registry struct {
	dependencies map[string]any
}

func NewRegistry() *Registry {
	return &Registry{
		dependencies: make(map[string]any),
	}
}

func (r *Registry) Provide(key string, dependency any) {
	r.dependencies[key] = dependency
}
func (r *Registry) Inject(key string) (dependency any) {
	dependency, ok := r.dependencies[key]
	if !ok {
		panic(fmt.Sprintf("dependency not found: %s", key))
	}
	return dependency
}
