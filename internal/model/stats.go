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
		switch visit.Role {
		case FieldRoleProperty:
			stats.Properties++
		case FieldRoleMultiField:
			stats.MultiFields++
		case FieldRoleRuntimeField:
			stats.RuntimeFields++
		}
	})
	stats.TotalFields = stats.Properties + stats.MultiFields + stats.RuntimeFields
	return stats
}
