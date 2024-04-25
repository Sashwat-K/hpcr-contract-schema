package hpcrcontractschema

import (
	_ "embed"
)

//go:embed certificate/hpse-contract-schema-1.0.56.json
var ContractSchema string
