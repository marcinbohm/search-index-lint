package rules

import "github.com/marcinbohm/search-index-lint/internal/model"

type Rule interface {
	Metadata() Metadata
	Check(ctx Context, corpus Corpus) ([]model.Finding, error)
}
