package diffrules

import (
	"strings"
	"testing"

	"github.com/marcinbohm/search-index-preflight/internal/diff"
	"github.com/marcinbohm/search-index-preflight/internal/model"
)

func TestDIF001EmitsFindingForFieldTypeChanged(t *testing.T) {
	findings := runDIF001(t, diff.Result{
		FieldChanges: []diff.FieldChange{
			fieldTypeChanged("status", model.FieldRoleProperty, "keyword", "long", "/properties/status", "/properties/status"),
		},
	})

	if len(findings) != 1 {
		t.Fatalf("expected one finding, got %#v", findings)
	}
	finding := findings[0]
	if finding.ID != "DIF001" {
		t.Fatalf("expected ID DIF001, got %q", finding.ID)
	}
	if finding.Severity != model.SeverityError {
		t.Fatalf("expected severity error, got %q", finding.Severity)
	}
	if finding.Confidence != model.ConfidenceHigh {
		t.Fatalf("expected confidence high, got %q", finding.Confidence)
	}
	if finding.Determinism != model.DeterminismDeterministic {
		t.Fatalf("expected deterministic finding, got %q", finding.Determinism)
	}
	if finding.Category != "schema-diff" {
		t.Fatalf("expected category schema-diff, got %q", finding.Category)
	}
	if finding.File != "mapping.json" {
		t.Fatalf("expected file mapping.json, got %q", finding.File)
	}
	if finding.JSONPointer != "/properties/status" {
		t.Fatalf("expected JSON pointer /properties/status, got %q", finding.JSONPointer)
	}
	for _, text := range []string{"status", "keyword", "long"} {
		if !strings.Contains(finding.Message, text) {
			t.Fatalf("expected message %q to contain %q", finding.Message, text)
		}
	}
	if finding.Remediation == "" {
		t.Fatal("expected remediation")
	}
	if finding.Fingerprint == "" {
		t.Fatal("expected fingerprint")
	}
}

func TestDIF001IgnoresFieldAdded(t *testing.T) {
	findings := runDIF001(t, diff.Result{
		FieldChanges: []diff.FieldChange{
			{
				Kind:     diff.ChangeFieldAdded,
				Resource: diff.ResourceID{Kind: diff.ResourceMapping, File: "mapping.json"},
				Field:    diff.FieldID{Path: "status", Role: model.FieldRoleProperty},
				After:    snapshot("status", model.FieldRoleProperty, "keyword", "/properties/status"),
			},
		},
	})

	if len(findings) != 0 {
		t.Fatalf("expected no findings, got %#v", findings)
	}
}

func TestDIF001IgnoresFieldRemoved(t *testing.T) {
	findings := runDIF001(t, diff.Result{
		FieldChanges: []diff.FieldChange{
			{
				Kind:     diff.ChangeFieldRemoved,
				Resource: diff.ResourceID{Kind: diff.ResourceMapping, File: "mapping.json"},
				Field:    diff.FieldID{Path: "status", Role: model.FieldRoleProperty},
				Before:   snapshot("status", model.FieldRoleProperty, "keyword", "/properties/status"),
			},
		},
	})

	if len(findings) != 0 {
		t.Fatalf("expected no findings, got %#v", findings)
	}
}

func TestDIF001UsesBeforePointerWhenAfterIsNil(t *testing.T) {
	findings := runDIF001(t, diff.Result{
		FieldChanges: []diff.FieldChange{
			{
				Kind:     diff.ChangeFieldTypeChanged,
				Resource: diff.ResourceID{Kind: diff.ResourceMapping, File: "mapping.json"},
				Field:    diff.FieldID{Path: "status", Role: model.FieldRoleProperty},
				Before:   snapshot("status", model.FieldRoleProperty, "keyword", "/properties/status"),
			},
		},
	})

	if len(findings) != 1 {
		t.Fatalf("expected one finding, got %#v", findings)
	}
	if findings[0].JSONPointer != "/properties/status" {
		t.Fatalf("expected before pointer, got %q", findings[0].JSONPointer)
	}
}

func TestDIF001WorksForMultiFieldTypeChanged(t *testing.T) {
	findings := runDIF001(t, diff.Result{
		FieldChanges: []diff.FieldChange{
			fieldTypeChanged("message.keyword", model.FieldRoleMultiField, "keyword", "wildcard", "/properties/message/fields/keyword", "/properties/message/fields/keyword"),
		},
	})

	if len(findings) != 1 {
		t.Fatalf("expected one finding, got %#v", findings)
	}
	if !strings.Contains(findings[0].Message, "message.keyword") {
		t.Fatalf("expected message to contain multi-field path, got %q", findings[0].Message)
	}
	if findings[0].JSONPointer != "/properties/message/fields/keyword" {
		t.Fatalf("expected multi-field pointer, got %q", findings[0].JSONPointer)
	}
}

func TestDIF001WorksForRuntimeFieldTypeChanged(t *testing.T) {
	findings := runDIF001(t, diff.Result{
		FieldChanges: []diff.FieldChange{
			fieldTypeChanged("day_of_week", model.FieldRoleRuntimeField, "keyword", "long", "/runtime/day_of_week", "/runtime/day_of_week"),
		},
	})

	if len(findings) != 1 {
		t.Fatalf("expected one finding, got %#v", findings)
	}
	if !strings.Contains(findings[0].Message, "day_of_week") {
		t.Fatalf("expected message to contain runtime field path, got %q", findings[0].Message)
	}
	if findings[0].JSONPointer != "/runtime/day_of_week" {
		t.Fatalf("expected runtime pointer, got %q", findings[0].JSONPointer)
	}
}

func runDIF001(t *testing.T, result diff.Result) []model.Finding {
	t.Helper()
	findings, err := NewDIF001().Check(Context{}, result)
	if err != nil {
		t.Fatalf("DIF001 returned error: %v", err)
	}
	return findings
}

func fieldTypeChanged(path string, role model.FieldRole, beforeType string, afterType string, beforePointer string, afterPointer string) diff.FieldChange {
	return diff.FieldChange{
		Kind:     diff.ChangeFieldTypeChanged,
		Resource: diff.ResourceID{Kind: diff.ResourceMapping, File: "mapping.json"},
		Field:    diff.FieldID{Path: path, Role: role},
		Before:   snapshot(path, role, beforeType, beforePointer),
		After:    snapshot(path, role, afterType, afterPointer),
	}
}

func snapshot(path string, role model.FieldRole, typ string, pointer string) *diff.FieldSnapshot {
	return &diff.FieldSnapshot{
		Path:        path,
		Role:        role,
		Type:        typ,
		JSONPointer: pointer,
	}
}
