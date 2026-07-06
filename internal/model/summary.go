package model

type Summary struct {
	FilesScanned  int `json:"files_scanned,omitempty"`
	FindingsTotal int `json:"findings_total,omitempty"`
	Critical      int `json:"critical,omitempty"`
	Error         int `json:"error,omitempty"`
	Warning       int `json:"warning,omitempty"`
	Info          int `json:"info,omitempty"`
	ExitCode      int `json:"exit_code,omitempty"`
}

type Tool struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type RunResult struct {
	SchemaVersion string       `json:"schema_version"`
	Tool          Tool         `json:"tool"`
	Summary       Summary      `json:"summary"`
	Findings      []Finding    `json:"findings"`
	Diagnostics   []Diagnostic `json:"diagnostics"`
}
