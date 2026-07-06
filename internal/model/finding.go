package model

type Finding struct {
	ID          string      `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Severity    Severity    `json:"severity,omitempty"`
	Confidence  Confidence  `json:"confidence,omitempty"`
	Category    string      `json:"category,omitempty"`
	Determinism Determinism `json:"determinism,omitempty"`
	File        string      `json:"file,omitempty"`
	JSONPointer string      `json:"json_pointer,omitempty"`
	Message     string      `json:"message,omitempty"`
	Remediation string      `json:"remediation,omitempty"`
	Fingerprint string      `json:"fingerprint,omitempty"`
}
