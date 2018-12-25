package compliance

import (
	"github.com/asaskevich/govalidator"
	"github.com/spn/go/address"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
	govalidator.CustomTypeTagMap.Set("spn_address", govalidator.CustomTypeValidator(isSpnAddress))
}

func isSpnAddress(i interface{}, context interface{}) bool {
	addr, ok := i.(string)

	if !ok {
		return false
	}

	_, _, err := address.Split(addr)

	if err == nil {
		return true
	}

	return false
}
