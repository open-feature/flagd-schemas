package flagd_definitions

import _ "embed"

//go:embed flagd-definitions.json
var FlagdDefinitions string

//go:embed targeting.json
var Targeting string
