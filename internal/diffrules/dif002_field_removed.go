package diffrules

import (
	"fmt"

	"github.com/marcinbohm/search-index-preflight/internal/diff"
	"github.com/marcinbohm/search-index-preflight/internal/model"
)

type dif002FieldRemoved struct{}

func NewDIF002() Rule {
	return dif002FieldRemoved{}
}

func (r dif002FieldRemoved) Metadata() Metadata {
	return Metadata{
		ID:          "DIF002",
		Name:        "field-removed",
		Category:    "schema-diff",
		Description: "Detects fields that were present in the base schema corpus but are missing from the current schema corpus.",
		Severity:    model.SeverityWarning,
		Confidence:  model.ConfidenceHigh,
		Determinism: model.DeterminismDeterministic,
	}
}

func (r dif002FieldRemoved) Check(ctx Context, result diff.Result) ([]model.Finding, error) {
	var findings []model.Finding
	for _, change := range result.FieldChanges {
		if change.Kind != diff.ChangeFieldRemoved {
			continue
		}
		findings = append(findings, r.finding(change))
	}
	return findings, nil
}

func (r dif002FieldRemoved) finding(change diff.FieldChange) model.Finding {
	metadata := r.Metadata()
	pointer := removedFieldPointer(change)

	return model.Finding{
		ID:          metadata.ID,
		Name:        metadata.Name,
		Severity:    metadata.Severity,
		Confidence:  metadata.Confidence,
		Category:    metadata.Category,
		Determinism: metadata.Determinism,
		File:        change.Resource.File,
		JSONPointer: pointer,
		Message:     fmt.Sprintf("Field %q was removed from the current schema.", change.Field.Path),
		Remediation: "Verify that producers, queries, dashboards, alerts, and downstream consumers no longer depend on this field. If the removal is intentional, coordinate it with index rollover or reindexing plans.",
		Fingerprint: fmt.Sprintf("%s:%s:%s:%s:%s:%s", metadata.ID, change.Resource.Kind, change.Resource.File, pointer, change.Field.Path, change.Field.Role),
	}
}

func removedFieldPointer(change diff.FieldChange) string {
	if change.Before != nil && change.Before.JSONPointer != "" {
		return change.Before.JSONPointer
	}
	return change.Resource.JSONPointer
}
