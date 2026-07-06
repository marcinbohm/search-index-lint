package model

type Corpus struct {
	Mappings           []Mapping
	IndexTemplates     []IndexTemplate
	ComponentTemplates []ComponentTemplate
	SampleDocuments    []RawDocument
	Diagnostics        []Diagnostic
}
