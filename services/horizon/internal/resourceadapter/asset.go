package resourceadapter

import (
	"context"

	"github.com/spn/go/xdr"
	. "github.com/spn/go/protocols/horizon"

)

func PopulateAsset(ctx context.Context, dest *Asset, asset xdr.Asset) error {
	return asset.Extract(&dest.Type, &dest.Code, &dest.Issuer)
}
