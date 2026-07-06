package model

type FieldStats struct {
	Properties    int
	MultiFields   int
	RuntimeFields int
	TotalFields   int
}

func CountFields(corpus Corpus) FieldStats {
	var stats FieldStats
	WalkFields(corpus, func(visit FieldVisit) {
		addFieldVisit(&stats, visit)
	})
	stats.TotalFields = stats.Properties + stats.MultiFields + stats.RuntimeFields
	return stats
}

func CountMappingFields(mapping Mapping) FieldStats {
	var stats FieldStats
	walkMappingFields(mapping, fieldVisitContext{
		origin:        FieldOriginMapping,
		mappingSource: mapping.Source,
	}, func(visit FieldVisit) {
		addFieldVisit(&stats, visit)
	})
	stats.TotalFields = stats.Properties + stats.MultiFields + stats.RuntimeFields
	return stats
}

func addFieldVisit(stats *FieldStats, visit FieldVisit) {
	switch visit.Role {
	case FieldRoleProperty:
		stats.Properties++
	case FieldRoleMultiField:
		stats.MultiFields++
	case FieldRoleRuntimeField:
		stats.RuntimeFields++
	}
}
