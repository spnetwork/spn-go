package helpers

import (
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/spn/go/address"
	"github.com/spn/go/amount"
	"github.com/spn/go/strkey"
)

func init() {
	govalidator.CustomTypeTagMap.Set("spn_accountid", govalidator.CustomTypeValidator(isSpnAccountID))
	govalidator.CustomTypeTagMap.Set("spn_seed", govalidator.CustomTypeValidator(isSpnSeed))
	govalidator.CustomTypeTagMap.Set("spn_asset_code", govalidator.CustomTypeValidator(isSpnAssetCode))
	govalidator.CustomTypeTagMap.Set("spn_address", govalidator.CustomTypeValidator(isSpnAddress))
	govalidator.CustomTypeTagMap.Set("spn_amount", govalidator.CustomTypeValidator(isSpnAmount))
	govalidator.CustomTypeTagMap.Set("spn_destination", govalidator.CustomTypeValidator(isSpnDestination))

}

func Validate(request Request, params ...interface{}) error {
	valid, err := govalidator.ValidateStruct(request)

	if !valid {
		fields := govalidator.ErrorsByField(err)
		for field, errorValue := range fields {
			switch {
			case errorValue == "non zero value required":
				return NewMissingParameter(field)
			case strings.HasSuffix(errorValue, "does not validate as spn_accountid"):
				return NewInvalidParameterError(field, "Account ID must start with `G` and contain 56 alphanum characters.")
			case strings.HasSuffix(errorValue, "does not validate as spn_seed"):
				return NewInvalidParameterError(field, "Account secret must start with `S` and contain 56 alphanum characters.")
			case strings.HasSuffix(errorValue, "does not validate as spn_asset_code"):
				return NewInvalidParameterError(field, "Asset code must be 1-12 alphanumeric characters.")
			case strings.HasSuffix(errorValue, "does not validate as spn_address"):
				return NewInvalidParameterError(field, "Spn address must be of form user*domain.com")
			case strings.HasSuffix(errorValue, "does not validate as spn_destination"):
				return NewInvalidParameterError(field, "Spn destination must be of form user*domain.com or start with `G` and contain 56 alphanum characters.")
			case strings.HasSuffix(errorValue, "does not validate as spn_amount"):
				return NewInvalidParameterError(field, "Amount must be positive and have up to 7 decimal places.")
			default:
				return NewInvalidParameterError(field, errorValue)
			}
		}
	}

	return request.Validate(params...)
}

// These are copied from support/config. Should we move them to /strkey maybe?
func isSpnAccountID(i interface{}, context interface{}) bool {
	enc, ok := i.(string)

	if !ok {
		return false
	}

	_, err := strkey.Decode(strkey.VersionByteAccountID, enc)

	if err == nil {
		return true
	}

	return false
}

func isSpnSeed(i interface{}, context interface{}) bool {
	enc, ok := i.(string)

	if !ok {
		return false
	}

	_, err := strkey.Decode(strkey.VersionByteSeed, enc)

	if err == nil {
		return true
	}

	return false
}

func isSpnAssetCode(i interface{}, context interface{}) bool {
	code, ok := i.(string)

	if !ok {
		return false
	}

	if !govalidator.IsByteLength(code, 1, 12) {
		return false
	}

	if !govalidator.IsAlphanumeric(code) {
		return false
	}

	return true
}

func isSpnAddress(i interface{}, context interface{}) bool {
	addr, ok := i.(string)

	if !ok {
		return false
	}

	_, _, err := address.Split(addr)
	if err != nil {
		return false
	}

	return true
}

func isSpnAmount(i interface{}, context interface{}) bool {
	am, ok := i.(string)

	if !ok {
		return false
	}

	_, err := amount.Parse(am)
	if err != nil {
		return false
	}

	return true
}

// isSpnDestination checks if `i` is either account public key or Spn address.
func isSpnDestination(i interface{}, context interface{}) bool {
	dest, ok := i.(string)

	if !ok {
		return false
	}

	_, err1 := strkey.Decode(strkey.VersionByteAccountID, dest)
	_, _, err2 := address.Split(dest)

	if err1 != nil && err2 != nil {
		return false
	}

	return true
}
