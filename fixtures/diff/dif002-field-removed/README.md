# DIF002 field removed

This fixture compares a base mapping that contains `legacy_id` with a current
mapping where that field is no longer present.

Expected:

- `search-index-preflight diff --base base --current current`
- exit code `0` with the default `--fail-on error`
- one warning `DIF002` finding
- exit code `1` when run with `--fail-on warning`

The data is synthetic and public-safe.
