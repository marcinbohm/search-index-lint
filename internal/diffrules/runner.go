package diffrules

import (
	"fmt"
)

func Run(ctx Context, registry *Registry, request RunRequest) (RunResult, error) {
	if registry == nil {
		return RunResult{}, fmt.Errorf("diff rule registry is nil")
	}

	var result RunResult
	for _, rule := range registry.List() {
		findings, err := rule.Check(ctx, request.Result)
		if err != nil {
			return result, fmt.Errorf("diff rule %s: %w", rule.Metadata().ID, err)
		}
		result.Findings = append(result.Findings, findings...)
	}
	return result, nil
}
