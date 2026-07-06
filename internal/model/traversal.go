package model

func WalkFields(corpus Corpus, visit func(FieldVisit)) {
	for _, mapping := range corpus.Mappings {
		walkMappingFields(mapping, fieldVisitContext{
			origin:        FieldOriginMapping,
			mappingSource: mapping.Source,
		}, visit)
	}
	for _, template := range corpus.IndexTemplates {
		if template.Template.Mappings == nil {
			continue
		}
		walkMappingFields(*template.Template.Mappings, fieldVisitContext{
			origin:            FieldOriginIndexTemplate,
			mappingSource:     template.Template.Mappings.Source,
			indexTemplateName: template.Name,
		}, visit)
	}
	for _, template := range corpus.ComponentTemplates {
		if template.Template.Mappings == nil {
			continue
		}
		walkMappingFields(*template.Template.Mappings, fieldVisitContext{
			origin:                FieldOriginComponentTemplate,
			mappingSource:         template.Template.Mappings.Source,
			componentTemplateName: template.Name,
		}, visit)
	}
}

func CollectFields(corpus Corpus) []FieldVisit {
	var visits []FieldVisit
	WalkFields(corpus, func(visit FieldVisit) {
		visits = append(visits, visit)
	})
	return visits
}

type fieldVisitContext struct {
	origin                FieldOrigin
	mappingSource         Source
	indexTemplateName     string
	componentTemplateName string
}

func walkMappingFields(mapping Mapping, context fieldVisitContext, visit func(FieldVisit)) {
	walkFields(mapping.Properties, context, FieldRoleProperty, visit)
	walkFields(mapping.RuntimeFields, context, FieldRoleRuntimeField, visit)
}

func walkFields(fields []Field, context fieldVisitContext, role FieldRole, visit func(FieldVisit)) {
	for _, field := range fields {
		visit(FieldVisit{
			Origin:                context.origin,
			Role:                  role,
			Source:                field.Source,
			MappingSource:         context.mappingSource,
			IndexTemplateName:     context.indexTemplateName,
			ComponentTemplateName: context.componentTemplateName,
			Field:                 field,
			Path:                  field.Path,
			JSONPointer:           field.JSONPointer,
		})
		walkFields(field.Properties, context, FieldRoleProperty, visit)
		walkFields(field.Fields, context, FieldRoleMultiField, visit)
	}
}
