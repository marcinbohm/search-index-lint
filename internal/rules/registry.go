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
	if err := validateMetadata(metadata); err != nil {
		return err
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

func validateMetadata(metadata Metadata) error {
	if metadata.ID == "" {
		return fmt.Errorf("rule ID is required")
	}
	if metadata.Name == "" {
		return fmt.Errorf("rule %q name is required", metadata.ID)
	}
	if metadata.Category == "" {
		return fmt.Errorf("rule %q category is required", metadata.ID)
	}
	if metadata.Severity == "" {
		return fmt.Errorf("rule %q severity is required", metadata.ID)
	}
	if metadata.Confidence == "" {
		return fmt.Errorf("rule %q confidence is required", metadata.ID)
	}
	if metadata.Determinism == "" {
		return fmt.Errorf("rule %q determinism is required", metadata.ID)
	}
	return nil
}
