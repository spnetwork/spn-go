package ethereum

import (
	"math/big"

	"github.com/spn/go/services/bifrost/common"
)

func (t Transaction) ValueToSpn() string {
	valueEth := new(big.Rat)
	valueEth.Quo(new(big.Rat).SetInt(t.ValueWei), weiInEth)
	return valueEth.FloatString(common.SpnAmountPrecision)
}
