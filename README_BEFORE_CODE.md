# SearchIndexLint

Catch Elasticsearch/OpenSearch mapping and template risks before production.

SearchIndexLint is a planned offline-first CLI and GitHub Action for linting Elasticsearch/OpenSearch mappings, component templates, index templates, dynamic templates, and sample documents.

Status: design-first, pre-code.

## Problem

Search schema changes can be syntactically valid and still operationally risky.

Examples:

- dynamic mappings causing mapping explosion
- field type conflicts across indices/templates
- broad dynamic templates
- dotted field collisions
- template priority conflicts
- bad `text` vs `keyword` choices
- missing analyzer/normalizer definitions
- sample documents that do not match declared mappings
- Elasticsearch/OpenSearch compatibility gaps

These failures often appear after indexing starts, when remediation may require rollover, reindexing, DLQ drain, producer fixes, or manual operator work.

## Goal

SearchIndexLint should provide pre-merge and pre-deploy feedback for schema-as-code repositories.

The MVP will be offline, CLI-first, fixture-driven, rule-based, explainable, safe for CI, and useful without cluster credentials.

## Non-goals

SearchIndexLint is not a dashboard, SaaS product, mapping generator, automatic fixer, cluster doctor, OpenSearch Dashboards plugin, replacement for staging tests, or replacement for operator judgment.

No live cluster mode in MVP.

## Planned CLI

```bash
search-index-lint lint --mapping mapping.json
search-index-lint lint --template index-template.json
search-index-lint lint --sample-docs samples.jsonl
search-index-lint lint ./schemas --format json
search-index-lint lint ./schemas --format markdown
search-index-lint lint ./schemas --format sarif
search-index-lint rules list
search-index-lint explain SIL001
```

## Initial MVP rules

- SIL001 total fields limit risk
- SIL002 root dynamic mapping enabled
- SIL003 dynamic template missing `match_mapping_type`
- SIL004 overbroad dynamic template
- SIL005 dynamic template shadowing
- SIL006 `path_match` object collision risk
- SIL007 dotted field collision
- SIL008 field type conflict
- SIL009 sample document conflicts with mapping
- SIL010 dynamic date/numeric detection risk
- SIL011 likely aggregatable field mapped as `text`
- SIL012 long `keyword` without `ignore_above`
- SIL013 `fielddata: true` on `text`
- SIL014 missing analyzer/normalizer definition
- SIL015 template priority conflict

## Safety note

SearchIndexLint detects risk. It does not prove a schema is safe.

A clean report does not replace tests, staging validation, load testing, observability, rollout planning, or operator review.

## Feedback

Feedback is welcome on the rule catalog, fixture examples, CLI contract, and false-positive risks.

Do not share confidential mappings, production logs, customer data, credentials, internal service names, or private cluster details in public issues.
