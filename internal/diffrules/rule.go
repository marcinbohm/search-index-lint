package diffrules

import (
	"github.com/marcinbohm/search-index-preflight/internal/diff"
	"github.com/marcinbohm/search-index-preflight/internal/model"
)

// Rule checks a semantic diff result and returns preflight findings.
type Rule interface {
	Metadata() Metadata
	Check(ctx Context, result diff.Result) ([]model.Finding, error)
}
