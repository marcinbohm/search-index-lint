# Offline migration/versioning concept

## Status

Planned, not implemented.

This is planned future work. No migration/versioning CLI commands are implemented today.

## Goal

SearchIndexPreflight may eventually validate Elasticsearch/OpenSearch schema evolution offline:

```text
versioned schema directories or migration manifests -> validate chain -> lint each state -> diff consecutive states -> report preflight risks
```

The goal is to help reviewers understand whether a schema change sequence is likely safe to merge or deploy, without applying anything to a cluster.

## Non-goals

This concept is not a cluster migration executor.

Non-goals:

- no apply
- no cluster writes
- no reindex execution
- no alias cutover execution
- no rollback execution
- no migration locks
- no migration state stored in Elasticsearch/OpenSearch
- no direct deployment orchestration

## How it builds on current foundations

The future layer should reuse existing SearchIndexPreflight foundations:

- input discovery
- JSON parser
- normalizer
- canonical `model.Corpus`
- static lint rules `SIL001` through `SIL003`
- public `diff` command behavior
- internal diff engine
- diff rules `DIF001` through `DIF003`
- console and JSON reports
- fixture-driven examples and regression tests

It should not reimplement parser, normalizer, lint, diff, or report foundations.

## Possible input models

These are exploratory models, not final contracts.

### Option A: Versioned directories

```text
schema-versions/
  001-initial/
    mapping.json
  002-add-customer-id/
    mapping.json
  003-remove-legacy-id/
    mapping.json
```

### Option B: Migration manifests

```text
migrations/
  001-initial.json
  002-add-customer-id.json
  003-remove-legacy-id.json
```

### Option C: Manifest plus schema snapshots

A future format may combine a small manifest with full schema snapshots for each state. This could preserve reviewer intent while keeping validation grounded in normalized schema corpora.

YAML may be considered later, but YAML input is not implemented today.

## Possible future commands

Command names are examples only and are not implemented:

```bash
search-index-preflight versions validate schema-versions/
search-index-preflight migrations validate migrations/
search-index-preflight migrations plan --base schema-versions/001 --current schema-versions/002
```

If a future command uses the word `migrations`, it must remain offline and preflight-only. It must not apply changes to Elasticsearch/OpenSearch.

## Example workflow

Conceptual output:

```text
Version chain: 001-initial -> 002-add-customer-id -> 003-remove-legacy-id

001-initial -> 002-add-customer-id
  info DIF003: customer_id added

002-add-customer-id -> 003-remove-legacy-id
  warning DIF002: legacy_id removed

Result:
  0 errors
  1 warning
  1 info
```

This output is illustrative. No migration/versioning report format exists yet.

## Safety model

The future layer should be advisory and preflight-only:

- fail on errors by default, consistent with current `--fail-on error`
- allow warning/info thresholds similar to the existing `--fail-on` behavior
- avoid cluster mutation
- avoid deployment orchestration
- produce human-readable and machine-readable preflight reports
- make operational steps explicit for reviewers without executing them

## Open questions

- Should command naming use `versions`, `migrations`, or another term?
- What manifest format should be supported?
- What version ID format should be required?
- Should gaps in numbering be errors or warnings?
- Should snapshots be required, or should operation-only manifests be allowed?
- Should the tool produce machine-readable migration plan JSON?
- How should index templates and component templates be represented?
- How should multiple indices/templates in one version be modeled?
- Should compatibility profiles influence version-chain validation later?

## Relationship to current lint/diff

Current implemented behavior remains unchanged:

- `lint` runs static offline checks and emits `SIL001`, `SIL002`, and `SIL003`
- `diff --base <path> --current <path>` compares two schema inputs and emits `DIF001`, `DIF002`, and `DIF003`

The migration/versioning concept would orchestrate those existing checks across ordered schema states. It is a future layer above current lint and diff behavior, not a replacement for them.
