package model

import (
	"reflect"
	"testing"
)

func TestCollectFieldsStandaloneMapping(t *testing.T) {
	source := Source{Path: "mapping.json", RelativePath: "mapping.json"}
	corpus := Corpus{
		Mappings: []Mapping{
			{
				Source: source,
				Properties: []Field{
					testField(source, "status", FieldRoleProperty),
					fieldWithMultiField(source, "message", "keyword"),
					fieldWithChild(source, "user", testField(source, "user.id", FieldRoleProperty)),
				},
				RuntimeFields: []Field{testField(source, "day", FieldRoleRuntimeField)},
			},
		},
	}

	visits := CollectFields(corpus)
	got := visitPathRoles(visits)
	want := []pathRole{
		{path: "status", role: FieldRoleProperty},
		{path: "message", role: FieldRoleProperty},
		{path: "message.keyword", role: FieldRoleMultiField},
		{path: "user", role: FieldRoleProperty},
		{path: "user.id", role: FieldRoleProperty},
		{path: "day", role: FieldRoleRuntimeField},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("CollectFields returned path/roles %#v, want %#v", got, want)
	}
	for _, visit := range visits {
		if visit.Origin != FieldOriginMapping {
			t.Fatalf("Origin = %q, want %q", visit.Origin, FieldOriginMapping)
		}
		if visit.Source.RelativePath != "mapping.json" {
			t.Fatalf("Source.RelativePath = %q, want mapping.json", visit.Source.RelativePath)
		}
	}
}

func TestCollectFieldsIndexTemplate(t *testing.T) {
	source := Source{Path: "index-template.json", RelativePath: "index-template.json"}
	corpus := Corpus{
		IndexTemplates: []IndexTemplate{
			{
				Source: source,
				Template: TemplateBody{
					Mappings: &Mapping{
						Source:     source,
						Properties: []Field{testField(source, "@timestamp", FieldRoleProperty)},
					},
				},
			},
		},
	}

	visits := CollectFields(corpus)
	if len(visits) != 1 {
		t.Fatalf("CollectFields returned %d visits, want 1", len(visits))
	}
	visit := visits[0]
	if visit.Origin != FieldOriginIndexTemplate {
		t.Fatalf("Origin = %q, want %q", visit.Origin, FieldOriginIndexTemplate)
	}
	if visit.Role != FieldRoleProperty {
		t.Fatalf("Role = %q, want %q", visit.Role, FieldRoleProperty)
	}
	if visit.Source.RelativePath != "index-template.json" {
		t.Fatalf("Source.RelativePath = %q, want index-template.json", visit.Source.RelativePath)
	}
	if visit.IndexTemplateName != "" {
		t.Fatalf("IndexTemplateName = %q, want empty", visit.IndexTemplateName)
	}
}

func TestCollectFieldsComponentTemplate(t *testing.T) {
	source := Source{Path: "component-template.json", RelativePath: "component-template.json"}
	corpus := Corpus{
		ComponentTemplates: []ComponentTemplate{
			{
				Source: source,
				Template: TemplateBody{
					Mappings: &Mapping{
						Source:     source,
						Properties: []Field{testField(source, "service.name", FieldRoleProperty)},
					},
				},
			},
		},
	}

	visits := CollectFields(corpus)
	if len(visits) != 1 {
		t.Fatalf("CollectFields returned %d visits, want 1", len(visits))
	}
	visit := visits[0]
	if visit.Origin != FieldOriginComponentTemplate {
		t.Fatalf("Origin = %q, want %q", visit.Origin, FieldOriginComponentTemplate)
	}
	if visit.Role != FieldRoleProperty {
		t.Fatalf("Role = %q, want %q", visit.Role, FieldRoleProperty)
	}
	if visit.ComponentTemplateName != "" {
		t.Fatalf("ComponentTemplateName = %q, want empty", visit.ComponentTemplateName)
	}
}

func TestCountFields(t *testing.T) {
	source := Source{Path: "mapping.json", RelativePath: "mapping.json"}
	corpus := Corpus{
		Mappings: []Mapping{
			{
				Source: source,
				Properties: []Field{
					testField(source, "status", FieldRoleProperty),
					fieldWithMultiField(source, "message", "keyword"),
					fieldWithChild(source, "user", testField(source, "user.id", FieldRoleProperty)),
				},
				RuntimeFields: []Field{testField(source, "day", FieldRoleRuntimeField)},
			},
		},
	}

	stats := CountFields(corpus)
	if stats.Properties != 4 {
		t.Fatalf("Properties = %d, want 4", stats.Properties)
	}
	if stats.MultiFields != 1 {
		t.Fatalf("MultiFields = %d, want 1", stats.MultiFields)
	}
	if stats.RuntimeFields != 1 {
		t.Fatalf("RuntimeFields = %d, want 1", stats.RuntimeFields)
	}
	if stats.TotalFields != 6 {
		t.Fatalf("TotalFields = %d, want 6", stats.TotalFields)
	}
}

type pathRole struct {
	path string
	role FieldRole
}

func visitPathRoles(visits []FieldVisit) []pathRole {
	values := make([]pathRole, 0, len(visits))
	for _, visit := range visits {
		values = append(values, pathRole{path: visit.Path, role: visit.Role})
	}
	return values
}

func testField(source Source, path string, role FieldRole) Field {
	name := path
	if lastDot := lastIndexByte(path, '.'); lastDot >= 0 {
		name = path[lastDot+1:]
	}
	field := Field{
		Name:        name,
		Path:        path,
		Source:      source,
		JSONPointer: "/" + path,
	}
	if role == FieldRoleMultiField {
		field.ParentPath = path[:lastIndexByte(path, '.')]
	}
	return field
}

func fieldWithMultiField(source Source, path, multiFieldName string) Field {
	value := testField(source, path, FieldRoleProperty)
	value.Fields = []Field{testField(source, path+"."+multiFieldName, FieldRoleMultiField)}
	return value
}

func fieldWithChild(source Source, path string, child Field) Field {
	value := testField(source, path, FieldRoleProperty)
	value.Properties = []Field{child}
	return value
}

func lastIndexByte(value string, target byte) int {
	for i := len(value) - 1; i >= 0; i-- {
		if value[i] == target {
			return i
		}
	}
	return -1
}
