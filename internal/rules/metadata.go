package rules

import "github.com/marcinbohm/search-index-lint/internal/model"

type Metadata struct {
	ID          string
	Name        string
	Category    string
	Description string
	Severity    model.Severity
	Confidence  model.Confidence
	Determinism model.Determinism
}
