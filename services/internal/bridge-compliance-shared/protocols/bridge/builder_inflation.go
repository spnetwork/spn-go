package bridge

import (
	b "github.com/spn/go/build"
	shared "github.com/spn/go/services/internal/bridge-compliance-shared"
	"github.com/spn/go/services/internal/bridge-compliance-shared/http/helpers"
)

// InflationOperationBody represents inflation operation
type InflationOperationBody struct {
	Source *string
}

// ToTransactionMutator returns go-spn-base TransactionMutator
func (op InflationOperationBody) ToTransactionMutator() b.TransactionMutator {
	var mutators []interface{}

	if op.Source != nil {
		mutators = append(mutators, b.SourceAccount{*op.Source})
	}

	return b.Inflation(mutators...)
}

// Validate validates if operation body is valid.
func (op InflationOperationBody) Validate() error {
	if op.Source != nil && !shared.IsValidAccountID(*op.Source) {
		return helpers.NewInvalidParameterError("source", "Source must be a public key (starting with `G`).")
	}

	return nil
}
