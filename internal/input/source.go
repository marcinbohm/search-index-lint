package input

import "github.com/marcinbohm/search-index-lint/internal/model"

type Source struct {
	Path         string
	RelativePath string
	Kind         model.DocumentKind
	Content      []byte
}
