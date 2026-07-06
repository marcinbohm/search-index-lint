package model

type DocumentKind string

const (
	DocumentKindUnknown           DocumentKind = "unknown"
	DocumentKindMapping           DocumentKind = "mapping"
	DocumentKindIndexTemplate     DocumentKind = "index_template"
	DocumentKindComponentTemplate DocumentKind = "component_template"
	DocumentKindSampleDocs        DocumentKind = "sample_docs"
)

type RawDocument struct {
	Kind        DocumentKind `json:"kind"`
	Source      Source       `json:"source"`
	Content     any          `json:"content,omitempty"`
	Diagnostics []Diagnostic `json:"diagnostics,omitempty"`
}
