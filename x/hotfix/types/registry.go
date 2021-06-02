package types

type (
	// Registry defines Hotfix registry
	Registry struct {
		m map[int64]*Hotfix
	}
)

// NewRegistry initializes Registry and returns the reference
func NewRegistry() *Registry {
	return &Registry{
		m: make(map[int64]*Hotfix),
	}
}

// WithHotfix sets the Hotfix for height
func (r *Registry) WithHotfix(v *Hotfix) *Registry {
	r.m[v.Height] = v
	return r
}

// Hotfix returns the handler of Hotfix for height
func (r *Registry) Hotfix(height int64) *Hotfix {
	return r.m[height]
}
