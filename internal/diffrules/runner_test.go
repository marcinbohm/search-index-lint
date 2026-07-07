package diffrules

import (
	"fmt"
	"testing"

	"github.com/marcinbohm/search-index-preflight/internal/diff"
	"github.com/marcinbohm/search-index-preflight/internal/model"
)

func TestBuiltinRegistryContainsOnlyDIF001(t *testing.T) {
	registry, err := BuiltinRegistry()
	if err != nil {
		t.Fatalf("BuiltinRegistry returned error: %v", err)
	}

	rules := registry.List()
	if len(rules) != 1 {
		t.Fatalf("expected one built-in diff rule, got %d", len(rules))
	}
	if rules[0].Metadata().ID != "DIF001" {
		t.Fatalf("expected DIF001, got %q", rules[0].Metadata().ID)
	}
}

func TestRegistryRejectsDuplicateIDs(t *testing.T) {
	_, err := NewRegistry(fakeRule{id: "DIF001"}, fakeRule{id: "DIF001"})
	if err == nil {
		t.Fatal("expected duplicate ID error")
	}
}

func TestRunExecutesBuiltinDIF001(t *testing.T) {
	registry, err := BuiltinRegistry()
	if err != nil {
		t.Fatalf("BuiltinRegistry returned error: %v", err)
	}

	result, err := Run(Context{}, registry, RunRequest{Result: diff.Result{
		FieldChanges: []diff.FieldChange{
			fieldTypeChanged("status", model.FieldRoleProperty, "keyword", "long", "/properties/status", "/properties/status"),
		},
	}})
	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	if len(result.Findings) != 1 {
		t.Fatalf("expected one finding, got %#v", result.Findings)
	}
	if result.Findings[0].ID != "DIF001" {
		t.Fatalf("expected DIF001 finding, got %q", result.Findings[0].ID)
	}
}

func TestDiffCompareToDIF001Integration(t *testing.T) {
	base := corpusWithMapping(property("status", "keyword"))
	current := corpusWithMapping(property("status", "long"))

	diffResult, err := diff.Compare(base, current)
	if err != nil {
		t.Fatalf("diff.Compare returned error: %v", err)
	}
	registry, err := BuiltinRegistry()
	if err != nil {
		t.Fatalf("BuiltinRegistry returned error: %v", err)
	}

	runResult, err := Run(Context{}, registry, RunRequest{Result: diffResult})
	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	if len(runResult.Findings) != 1 {
		t.Fatalf("expected one finding, got %#v", runResult.Findings)
	}
	if runResult.Findings[0].ID != "DIF001" {
		t.Fatalf("expected DIF001, got %q", runResult.Findings[0].ID)
	}
}

type fakeRule struct {
	id string
}

func (r fakeRule) Metadata() Metadata {
	return Metadata{
		ID:          r.id,
		Name:        "fake",
		Category:    "test",
		Description: "fake test rule",
		Severity:    model.SeverityWarning,
		Confidence:  model.ConfidenceHigh,
		Determinism: model.DeterminismDeterministic,
	}
}

func (r fakeRule) Check(ctx Context, result diff.Result) ([]model.Finding, error) {
	return nil, nil
}

func corpusWithMapping(fields ...model.Field) model.Corpus {
	return model.Corpus{
		Mappings: []model.Mapping{
			{
				Source:      model.Source{Path: "mapping.json", RelativePath: "mapping.json"},
				Properties:  fields,
				JSONPointer: "",
			},
		},
	}
}

func property(path string, typ string) model.Field {
	return model.Field{
		Name:        path,
		Path:        path,
		Type:        typ,
		Source:      model.Source{Path: "mapping.json", RelativePath: "mapping.json"},
		JSONPointer: fmt.Sprintf("/properties/%s", path),
	}
}
