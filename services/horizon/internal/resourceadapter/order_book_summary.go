package resourceadapter

import (
	"context"

	"github.com/spn/go/services/horizon/internal/db2/core"
	"github.com/spn/go/xdr"
	. "github.com/spn/go/protocols/horizon"
)

func PopulateOrderBookSummary(
	ctx context.Context,
	dest *OrderBookSummary,
	selling xdr.Asset,
	buying xdr.Asset,
	row core.OrderBookSummary,
) error {

	err := PopulateAsset(ctx, &dest.Selling, selling)
	if err != nil {
		return err
	}
	err = PopulateAsset(ctx, &dest.Buying, buying)
	if err != nil {
		return err
	}

	populatePriceLevels(&dest.Bids, row.Bids())
	populatePriceLevels(&dest.Asks, row.Asks())

	return nil
}

func populatePriceLevels(destp *[]PriceLevel, rows []core.OrderBookSummaryPriceLevel) {
	*destp = make([]PriceLevel, len(rows))
	dest := *destp

	for i, row := range rows {
		dest[i] = PriceLevel{
			Price:  row.PriceAsString(),
			Amount: row.AmountAsString(),
			PriceR: Price{
				N: row.Pricen,
				D: row.Priced,
			},
		}
	}
}
