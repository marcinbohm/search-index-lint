package model

import "testing"

func TestSeverityOrdering(t *testing.T) {
	tests := []struct {
		name  string
		left  Severity
		right Severity
		want  bool
	}{
		{name: "warning at least info", left: SeverityWarning, right: SeverityInfo, want: true},
		{name: "error at least warning", left: SeverityError, right: SeverityWarning, want: true},
		{name: "critical at least error", left: SeverityCritical, right: SeverityError, want: true},
		{name: "info below critical", left: SeverityInfo, right: SeverityCritical, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.left.AtLeast(tt.right); got != tt.want {
				t.Fatalf("%s.AtLeast(%s) = %v, want %v", tt.left, tt.right, got, tt.want)
			}
		})
	}
}

func TestParseSeverity(t *testing.T) {
	tests := []struct {
		input string
		want  Severity
	}{
		{input: "info", want: SeverityInfo},
		{input: " WARNING ", want: SeverityWarning},
		{input: "error", want: SeverityError},
		{input: "critical", want: SeverityCritical},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := ParseSeverity(tt.input)
			if err != nil {
				t.Fatalf("ParseSeverity(%q) returned error: %v", tt.input, err)
			}
			if got != tt.want {
				t.Fatalf("ParseSeverity(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestParseSeverityRejectsUnknownSeverity(t *testing.T) {
	if _, err := ParseSeverity("fatal"); err == nil {
		t.Fatal("ParseSeverity(\"fatal\") returned nil error")
	}
}
