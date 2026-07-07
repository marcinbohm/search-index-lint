# ADR 0004: Offline migration/versioning layer

## Status

Accepted.

## Date

2026-07-07

## Context

SearchIndexPreflight already supports offline static lint checks, a public experimental diff between base/current schema inputs, and basic field-level diff rules:

- `DIF001` field type changed
- `DIF002` field removed
- `DIF003` field added

Teams maintaining Elasticsearch/OpenSearch schemas often need more than a one-off diff. They need ordered schema versions, repeatable review of schema evolution, CI validation before merge, human-readable migration/preflight reports, and confidence that schema version changes do not introduce obvious breakage.

This is similar in spirit to Liquibase/Flyway-style schema evolution, but Elasticsearch/OpenSearch mappings, templates, aliases, rollovers, and reindexing workflows are different from relational database migrations.

## Decision

SearchIndexPreflight may add a future offline migration/versioning layer.

This layer is a helper and preflight layer only. It may validate and analyze versioned schema files or migration manifests, but it must not apply changes to Elasticsearch/OpenSearch clusters.

Allowed future capabilities:

- validate versioned schema directories
- validate migration manifests
- compare consecutive schema versions
- run existing lint rules against each version
- run existing diff rules between versions
- detect ordering and chain problems
- detect duplicate migration or version IDs
- report breaking-risk changes
- generate CI-friendly migration/versioning reports
- generate human-readable migration checklists
- help reviewers understand schema evolution before merge

## Non-Goals

The migration/versioning layer must not become a cluster migration executor.

Non-goals:

- no cluster writes
- no `migrate apply`
- no alias cutover execution
- no reindex execution
- no rollback execution
- no migration locks
- no migration state stored in Elasticsearch/OpenSearch
- no direct deployment orchestration
- no attempt to replace cluster/admin tooling

## Rationale

An offline migration/versioning layer fits the current preflight-only product direction. It can reuse the parser, normalizer, static lint rules, diff engine, diff rules, and report foundations already present in the project.

This gives teams a stronger CI workflow for schema-as-code repositories without turning SearchIndexPreflight into an operational tool that mutates clusters. Keeping the feature offline avoids high-risk failure modes around partial applies, alias cutovers, reindexing, rollback, cluster credentials, and live-state drift.

## Consequences

Positive consequences:

- clearer future direction beyond one-off lint and diff checks
- useful CI workflow for reviewing schema evolution
- stronger differentiation from a simple mapping linter
- no need for cluster credentials
- lower operational risk than a cluster-side migration executor

Tradeoffs:

- cannot guarantee live-cluster state
- cannot verify actual aliases, index templates, or component templates in a cluster
- cannot apply or rollback changes
- users must still execute operational migration steps themselves
- future command names must avoid implying cluster-side execution

## Naming Note

Final command names are not decided in this ADR.

Possible future command shapes include:

```bash
search-index-preflight versions validate
search-index-preflight migrations validate
search-index-preflight migrations plan
```

These commands are not implemented, and the names are not final. If `migrations` is used, the docs and help text must emphasize offline-only behavior and must not imply cluster-side apply, rollback, alias cutover, reindexing, or state management.
