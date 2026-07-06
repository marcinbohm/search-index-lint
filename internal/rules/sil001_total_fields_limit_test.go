package rules

import (
	"strconv"
	"testing"

	"github.com/marcinbohm/search-index-lint/internal/model"
)

func TestSIL001NoFindingUnderThreshold(t *testing.T) {
	findings := runSIL001(t, model.Mapping{
		Source:     testSource("mapping.json"),
		Properties: testFields(10),
	})
	if len(findings) != 0 {
		t.Fatalf("SIL001 returned findings %#v, want none", findings)
	}
}

func TestSIL001WarningNearDefaultLimit(t *testing.T) {
	findings := runSIL001(t, model.Mapping{
		Source:     testSource("mapping.json"),
		Properties: testFields(800),
	})
	requireSIL001Finding(t, findings, model.SeverityWarning)
}

func TestSIL001ErrorAtDefaultLimit(t *testing.T) {
	findings := runSIL001(t, model.Mapping{
		Source:     testSource("mapping.json"),
		Properties: testFields(1000),
	})
	requireSIL001Finding(t, findings, model.SeverityError)
}

func TestSIL001CountsMultiFields(t *testing.T) {
	fields := testFields(799)
	fields[0].Fields = []model.Field{
		{
			Name:        "keyword",
			Path:        fields[0].Path + ".keyword",
			ParentPath:  fields[0].Path,
			Source:      testSource("mapping.json"),
			JSONPointer: fields[0].JSONPointer + "/fields/keyword",
			Type:        "keyword",
		},
	}

	findings := runSIL001(t, model.Mapping{
		Source:     testSource("mapping.json"),
		Properties: fields,
	})
	requireSIL001Finding(t, findings, model.SeverityWarning)
}

func TestSIL001CountsRuntimeFields(t *testing.T) {
	findings := runSIL001(t, model.Mapping{
		Source:        testSource("mapping.json"),
		Properties:    testFields(799),
		RuntimeFields: testFields(1),
	})
	requireSIL001Finding(t, findings, model.SeverityWarning)
}

func runSIL001(t *testing.T, mapping model.Mapping) []model.Finding {
	t.Helper()
	findings, err := NewSIL001().Check(Context{}, model.Corpus{Mappings: []model.Mapping{mapping}})
	if err != nil {
		t.Fatalf("SIL001 returned error: %v", err)
	}
	return findings
}

func requireSIL001Finding(t *testing.T, findings []model.Finding, severity model.Severity) {
	t.Helper()
	if len(findings) != 1 {
		t.Fatalf("SIL001 returned %d findings, want 1: %#v", len(findings), findings)
	}
	finding := findings[0]
	if finding.ID != "SIL001" {
		t.Fatalf("finding ID = %q, want SIL001", finding.ID)
	}
	if finding.Severity != severity {
		t.Fatalf("finding severity = %q, want %q", finding.Severity, severity)
	}
	if finding.File != "mapping.json" {
		t.Fatalf("finding file = %q, want mapping.json", finding.File)
	}
	if finding.JSONPointer != "/" {
		t.Fatalf("finding JSON pointer = %q, want /", finding.JSONPointer)
	}
}

func testFields(count int) []model.Field {
	fields := make([]model.Field, 0, count)
	source := testSource("mapping.json")
	for i := 0; i < count; i++ {
		name := "field_" + strconv.Itoa(i)
		fields = append(fields, model.Field{
			Name:        name,
			Path:        name,
			Source:      source,
			JSONPointer: "/properties/" + name,
			Type:        "keyword",
		})
	}
	return fields
}

func testSource(path string) model.Source {
	return model.Source{Path: path, RelativePath: path}
}
