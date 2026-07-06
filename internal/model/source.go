package model

type Source struct {
	Path         string `json:"path,omitempty"`
	RelativePath string `json:"relative_path,omitempty"`
}
