---
title: "chore: add {{ env.TAG }} to schema store"
---

JSON schema released with tag {{ env.TAG }}. 
This needs to be added to the schema store catalog.  

[Schema store repo](https://github.com/SchemaStore/schemastore)

Extend the catalog entry found in `src/api/json/catalog.json` to include the new version:

```diff
{
    "name": "flagd flag configuration",
    "description": "Flag configuration for the OpenFeature spec compliant flagd provider",
    "fileMatch": [
    "flagd.json",
    "flagd.yaml",
    "flagd.yml",
    "*.flagd.json",
    "*.flagd.yaml",
    "*.flagd.yml"
    ],
    "url": "https://raw.githubusercontent.com/open-feature/schemas/main/json/flagd-definitions.json",
    "versions": {
    "0.1.1": "https://raw.githubusercontent.com/open-feature/schemas/json/json-schema-v0.1.1/json/flagd-definitions.json",
+   "X.X.X": "https://raw.githubusercontent.com/open-feature/schemas/json/json-schema-vX.X.X/json/flagd-definitions.json"
}
```