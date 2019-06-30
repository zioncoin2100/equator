---
title: Data
---

Each account in Zion network can contain multiple key/value pairs associated with it. Equator can be used to retrieve value of each data key.

When equator returns information about a single account data key it uses the following format:

## Attributes

| Attribute | Type | | 
| --- | --- | --- |
| value | base64-encoded string | The base64-encoded value for the key |

## Example

```json
{
  "value": "MTAw"
}
```
