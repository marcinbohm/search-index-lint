## Summary

Describe the change.

## Type of change

- [ ] Implementation
- [ ] Rule
- [ ] Parser/normalizer
- [ ] Reporter
- [ ] CLI UX
- [ ] Fixture
- [ ] Documentation
- [ ] Test-only
- [ ] Refactor
- [ ] CI/release
- [ ] Security
- [ ] Other

## Scope check

- [ ] This change fits the current milestone
- [ ] This does not add live cluster mode
- [ ] This does not add cluster write behavior
- [ ] This does not add auto-fix
- [ ] This does not add UI/dashboard/SaaS scope
- [ ] This does not include company/private data

## Tests

Commands run:

```bash
go test ./...
```

Additional tests:

- [ ] Unit tests added/updated
- [ ] Fixture tests added/updated
- [ ] Golden files added/updated
- [ ] CLI integration tests added/updated
- [ ] Compatibility tests added/updated
- [ ] Not applicable

## Fixtures

- [ ] Positive fixture added
- [ ] Negative fixture added
- [ ] Suppression fixture added
- [ ] Expected JSON updated
- [ ] Fixture README added/updated
- [ ] No fixture changes
- [ ] Fixture data is synthetic and public-safe

## Documentation

- [ ] README updated
- [ ] CLI contract updated
- [ ] Rule catalog updated
- [ ] Fixtures doc updated
- [ ] Architecture doc updated
- [ ] Contributing doc updated
- [ ] No docs needed

## Compatibility

Applies to:

- [ ] Elasticsearch
- [ ] OpenSearch
- [ ] Both
- [ ] Not applicable

Version impact:

- [ ] No version-specific behavior
- [ ] Version-specific behavior documented
- [ ] Compatibility fixtures added
- [ ] TBD / needs review

## Rule impact

For rule changes:

- Rule ID(s):
- Default severity:
- Confidence:
- Determinism:
- False-positive risk:
- New suppressions needed:
- Baseline impact:

Checklist:

- [ ] Rule metadata complete
- [ ] Rule docs updated
- [ ] `rules list` output updated
- [ ] `explain` output updated
- [ ] Remediation is cautious
- [ ] Heuristic behavior is not failing CI by default unless approved

## Reviewer checklist

- [ ] Architecture boundaries preserved
- [ ] CLI contract preserved or docs updated
- [ ] Report output deterministic
- [ ] Error handling is user-friendly
- [ ] No private data
- [ ] No hidden network calls
- [ ] Tests are meaningful
- [ ] Golden diffs are intentional
- [ ] Docs match behavior
- [ ] Future maintainability acceptable
