package types

type (
	Registry struct {
		m map[int64]*Hotfix
	}
)

func NewRegistry() *Registry {
	return &Registry{
		m: make(map[int64]*Hotfix),
	}
}

func (r *Registry) WithHotfix(v *Hotfix) *Registry {
	r.m[v.Height] = v
	return r
}

func (r *Registry) Hotfix(height int64) *Hotfix {
	return r.m[height]
}
