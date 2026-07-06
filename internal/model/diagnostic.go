package model

type Diagnostic struct {
	Severity Severity `json:"severity,omitempty"`
	File     string   `json:"file,omitempty"`
	Line     int      `json:"line,omitempty"`
	Message  string   `json:"message,omitempty"`
}
