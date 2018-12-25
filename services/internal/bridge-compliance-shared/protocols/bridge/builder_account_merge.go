package bridge

import (
	b "github.com/spn/go/build"
	shared "github.com/spn/go/services/internal/bridge-compliance-shared"
	"github.com/spn/go/services/internal/bridge-compliance-shared/http/helpers"
)

// AccountMergeOperationBody represents account_merge operation
type AccountMergeOperationBody struct {
	Source      *string
	Destination string
}

// ToTransactionMutator returns spn/go TransactionMutator
func (op AccountMergeOperationBody) ToTransactionMutator() b.TransactionMutator {
	mutators := []interface{}{b.Destination{op.Destination}}

	if op.Source != nil {
		mutators = append(mutators, b.SourceAccount{*op.Source})
	}

	return b.AccountMerge(mutators...)
}

// Validate validates if operation body is valid.
func (op AccountMergeOperationBody) Validate() error {
	if !shared.IsValidAccountID(op.Destination) {
		return helpers.NewInvalidParameterError("destination", "Destination must start with `G`.")
	}

	if op.Source != nil && !shared.IsValidAccountID(*op.Source) {
		return helpers.NewInvalidParameterError("source", "Source must start with `G`.")
	}

	return nil
}
