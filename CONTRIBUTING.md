## Guidelines

Welcome!

Please ensure you have read the [a code of conduct](https://github.com/open-feature/.github/blob/main/CODE_OF_CONDUCT.md) before contributing.  
TLDR: be respectful.


## Developing

### flagd JSON schema

The flagd JSON schema defines the json structure used within flagd for flag configurations.   
To contribute to the schema make changes in `json/flagd-definitions.yaml`, these changes can then be reflected in `json/flagd-definitions.json` by running `make gen-schema-json`.  
Unit testing for the schema is managed through the `json/test/negative` (in which all json files must fail validation against the schema) and `json/test/positive` (in which all json files must pass)
