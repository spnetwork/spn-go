// Package meta provides helpers for processing the metadata that is produced by
// spn-core while processing transactions.
package meta

import "github.com/spn/go/xdr"

// Bundle represents all of the metadata emitted from the application of a single
// spn transaction; Both fee meta and result meta is included.
type Bundle struct {
	FeeMeta         xdr.LedgerEntryChanges
	TransactionMeta xdr.TransactionMeta
}
