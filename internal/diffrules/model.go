package diffrules

import (
	"github.com/marcinbohm/search-index-preflight/internal/diff"
	"github.com/marcinbohm/search-index-preflight/internal/model"
)

// Metadata describes a diff-aware rule.
type Metadata struct {
	ID          string
	Name        string
	Category    string
	Description string
	Severity    model.Severity
	Confidence  model.Confidence
	Determinism model.Determinism
}

// Context carries shared diff-rule execution context.
type Context struct{}

type RunRequest struct {
	Result diff.Result
}

type RunResult struct {
	Findings []model.Finding
}
