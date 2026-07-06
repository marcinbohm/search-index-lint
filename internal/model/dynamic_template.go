package model

type DynamicTemplate struct {
	Name                string
	Source              Source
	JSONPointer         string
	Match               string
	Unmatch             string
	PathMatch           string
	PathUnmatch         string
	MatchMappingType    string
	HasMatchMappingType bool
	Mapping             map[string]any
}
