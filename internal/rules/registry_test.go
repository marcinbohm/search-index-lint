package rules

import (
	"testing"

	"github.com/marcinbohm/search-index-lint/internal/model"
)

type testRule struct {
	metadata Metadata
}

func (r testRule) Metadata() Metadata {
	return r.metadata
}

func (r testRule) Check(ctx Context, corpus Corpus) ([]model.Finding, error) {
	return nil, nil
}

func TestRegistryRejectsDuplicateRuleID(t *testing.T) {
	_, err := NewRegistry(
		testRule{metadata: Metadata{ID: "SIL001", Name: "first"}},
		testRule{metadata: Metadata{ID: "SIL001", Name: "second"}},
	)
	if err == nil {
		t.Fatal("NewRegistry returned nil error for duplicate rule ID")
	}
}

func TestRegistryListSortedByID(t *testing.T) {
	registry, err := NewRegistry(
		testRule{metadata: Metadata{ID: "SIL010", Name: "ten"}},
		testRule{metadata: Metadata{ID: "SIL002", Name: "two"}},
		testRule{metadata: Metadata{ID: "SIL001", Name: "one"}},
	)
	if err != nil {
		t.Fatalf("NewRegistry returned error: %v", err)
	}

	gotRules := registry.List()
	got := make([]string, 0, len(gotRules))
	for _, rule := range gotRules {
		got = append(got, rule.Metadata().ID)
	}

	want := []string{"SIL001", "SIL002", "SIL010"}
	if len(got) != len(want) {
		t.Fatalf("List returned %d rules, want %d: %v", len(got), len(want), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("List returned IDs %v, want %v", got, want)
		}
	}
}
