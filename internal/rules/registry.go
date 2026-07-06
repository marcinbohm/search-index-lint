package rules

import (
	"fmt"
	"sort"
)

type Registry struct {
	rules map[string]Rule
}

func NewRegistry(rules ...Rule) (*Registry, error) {
	registry := &Registry{rules: map[string]Rule{}}
	for _, rule := range rules {
		if err := registry.Register(rule); err != nil {
			return nil, err
		}
	}
	return registry, nil
}

func (r *Registry) Register(rule Rule) error {
	if rule == nil {
		return fmt.Errorf("rule is nil")
	}
	metadata := rule.Metadata()
	if metadata.ID == "" {
		return fmt.Errorf("rule ID is required")
	}
	if _, exists := r.rules[metadata.ID]; exists {
		return fmt.Errorf("duplicate rule ID %q", metadata.ID)
	}
	r.rules[metadata.ID] = rule
	return nil
}

func (r *Registry) Get(id string) (Rule, bool) {
	rule, ok := r.rules[id]
	return rule, ok
}

func (r *Registry) List() []Rule {
	list := make([]Rule, 0, len(r.rules))
	for _, rule := range r.rules {
		list = append(list, rule)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Metadata().ID < list[j].Metadata().ID
	})
	return list
}
