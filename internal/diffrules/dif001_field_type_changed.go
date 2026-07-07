package diffrules

import (
	"fmt"

	"github.com/marcinbohm/search-index-preflight/internal/diff"
	"github.com/marcinbohm/search-index-preflight/internal/model"
)

type dif001FieldTypeChanged struct{}

func NewDIF001() Rule {
	return dif001FieldTypeChanged{}
}

func (r dif001FieldTypeChanged) Metadata() Metadata {
	return Metadata{
		ID:          "DIF001",
		Name:        "field-type-changed",
		Category:    "schema-diff",
		Description: "Detects fields whose mapped type changed between base and current schema corpora.",
		Severity:    model.SeverityError,
		Confidence:  model.ConfidenceHigh,
		Determinism: model.DeterminismDeterministic,
	}
}

func (r dif001FieldTypeChanged) Check(ctx Context, result diff.Result) ([]model.Finding, error) {
	var findings []model.Finding
	for _, change := range result.FieldChanges {
		if change.Kind != diff.ChangeFieldTypeChanged {
			continue
		}
		findings = append(findings, r.finding(change))
	}
	return findings, nil
}

func (r dif001FieldTypeChanged) finding(change diff.FieldChange) model.Finding {
	metadata := r.Metadata()
	beforeType := snapshotType(change.Before)
	afterType := snapshotType(change.After)
	pointer := findingPointer(change)

	return model.Finding{
		ID:          metadata.ID,
		Name:        metadata.Name,
		Severity:    metadata.Severity,
		Confidence:  metadata.Confidence,
		Category:    metadata.Category,
		Determinism: metadata.Determinism,
		File:        change.Resource.File,
		JSONPointer: pointer,
		Message:     fmt.Sprintf("Field %q changed type from %q to %q.", change.Field.Path, beforeType, afterType),
		Remediation: "Changing a mapped field type usually requires creating a new index and reindexing or rolling over with a compatible migration plan. Verify whether existing indices and queries depend on the previous type.",
		Fingerprint: fmt.Sprintf("%s:%s:%s:%s:%s:%s", metadata.ID, change.Resource.Kind, change.Resource.File, pointer, change.Field.Path, change.Field.Role),
	}
}

func findingPointer(change diff.FieldChange) string {
	if change.After != nil && change.After.JSONPointer != "" {
		return change.After.JSONPointer
	}
	if change.Before != nil && change.Before.JSONPointer != "" {
		return change.Before.JSONPointer
	}
	return change.Resource.JSONPointer
}

func snapshotType(snapshot *diff.FieldSnapshot) string {
	if snapshot == nil {
		return ""
	}
	return snapshot.Type
}
