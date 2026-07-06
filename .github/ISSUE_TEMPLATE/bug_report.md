---
name: Bug report
about: Report incorrect behavior in SearchIndexLint
title: "bug: "
labels: ["type: bug"]
assignees: ""
---

## Summary

Describe the bug.

## Version

SearchIndexLint version:

```text
TBD
```

Operating system:

```text
TBD
```

## Command

```bash
search-index-lint lint ...
```

## Expected behavior

What did you expect to happen?

## Actual behavior

What happened instead?

## Minimal reproduction

Please provide the smallest synthetic example possible.

Do not include production mappings, logs, customer data, credentials, internal service names, or confidential index patterns.

### Mapping/template

```json
{
}
```

### Sample docs, if relevant

```jsonl
{"example":"value"}
```

### Config, if relevant

```yaml
version: 1
```

## Output

```text
Paste SearchIndexLint output here.
```

## Is this a false positive?

- [ ] Yes
- [ ] No
- [ ] Not sure

## Dialect

- [ ] Elasticsearch
- [ ] OpenSearch
- [ ] Unknown

Version:

```text
TBD
```

## Privacy check

- [ ] I removed production data
- [ ] I removed credentials/tokens
- [ ] I removed customer data
- [ ] I removed internal service names
- [ ] I reduced the example to a minimal synthetic reproduction
