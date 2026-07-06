# SIL003 Missing Match Mapping Type Fixtures

These fixtures exercise `SIL003` (`dynamic-template-missing-match-mapping-type`), which detects dynamic templates that omit `match_mapping_type`.

Why it matters:

- a dynamic template without `match_mapping_type` may apply to more inferred field types than intended
- broad matching can create mapping growth, type compatibility, or query behavior surprises
- broad dynamic templates may still be intentional for flexible schemas or controlled ingestion paths

This rule is heuristic. It reports a warning and does not mean the dynamic template is invalid.

Current scope:

- checks normalized dynamic templates for missing `match_mapping_type`
- does not validate whether a present `match_mapping_type` value is good or compatible
- does not estimate dynamic field expansion
- does not analyze dynamic template ordering or shadowing
- does not compose component templates

Fixture cases:

- `mapping-missing-match-mapping-type.json`: emits one `SIL003` warning and exits `0` with the default `--fail-on error`.
- `mapping-with-match-mapping-type.json`: emits no `SIL003` finding.

Remediation guidance:

- add `match_mapping_type` when the template is intended for a specific detected field type
- document why broad matching is intentional when omitting `match_mapping_type`

Privacy note: these fixtures are fully synthetic and contain no private, customer, company, or production data.
