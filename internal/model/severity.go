package model

import (
	"fmt"
	"strings"
)

type Severity string

const (
	SeverityInfo     Severity = "info"
	SeverityWarning  Severity = "warning"
	SeverityError    Severity = "error"
	SeverityCritical Severity = "critical"
)

var severityRank = map[Severity]int{
	SeverityInfo:     0,
	SeverityWarning:  1,
	SeverityError:    2,
	SeverityCritical: 3,
}

func ParseSeverity(value string) (Severity, error) {
	severity := Severity(strings.ToLower(strings.TrimSpace(value)))
	if _, ok := severityRank[severity]; !ok {
		return "", fmt.Errorf("unknown severity %q", value)
	}
	return severity, nil
}

func (s Severity) Rank() int {
	rank, ok := severityRank[s]
	if !ok {
		return -1
	}
	return rank
}

func (s Severity) AtLeast(other Severity) bool {
	return s.Rank() >= other.Rank()
}

type Confidence string

const (
	ConfidenceLow    Confidence = "low"
	ConfidenceMedium Confidence = "medium"
	ConfidenceHigh   Confidence = "high"
)

type Determinism string

const (
	DeterminismDeterministic          Determinism = "deterministic"
	DeterminismHeuristic              Determinism = "heuristic"
	DeterminismClusterContextRequired Determinism = "cluster-context-required"
)
