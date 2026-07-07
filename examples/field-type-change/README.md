# Field type change

This example shows a risky schema change where `status` changes from `keyword` to `long`.

## Run

```bash
go run ./cmd/search-index-preflight diff \
  --base examples/field-type-change/base \
  --current examples/field-type-change/current
```

## Expected finding

```text
error DIF001: mapping.json#/properties/status: Field "status" changed type from "keyword" to "long".
```

## Why this matters

Changing a mapped field type usually requires a new index, rollover, reindexing, or another explicit migration plan. Existing queries and dashboards may also depend on the previous type.
