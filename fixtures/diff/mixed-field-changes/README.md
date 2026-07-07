# Mixed field changes

This fixture emits all current public diff rules in one run.

Base:

- `status` is `keyword`
- `legacy_id` exists

Current:

- `status` is `long`
- `legacy_id` is removed
- `customer_id` is added

Expected:

- `DIF001` error for `status` type change
- `DIF002` warning for `legacy_id` removal
- `DIF003` info for `customer_id` addition
- default exit code `1`
