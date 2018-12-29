package logic

import (
	"fmt"
	"reflect"

	"github.com/rompi/tax-calc/app/model"
	"github.com/rompi/tax-calc/constant"
)

func Validate(o *model.Object) []string {
	var es []string
	if o.Name == "" {
		es = append(es, fmt.Sprintf("Name should not be empty"))
	}

	if !greaterThanEqualZero(o.Price) {
		es = append(es, fmt.Sprintf("Price must be greater than zero"))
	}

	if !validTaxCode(o.TaxCode) {
		es = append(es, fmt.Sprintf("Invalid Tax Code"))
	}

	return es
}

func greaterThanEqualZero(i interface{}) bool {
	kind := reflect.ValueOf(i).Kind().String()
	switch kind {
	case "float64":
		return i.(float64) >= float64(0)
	case "int":
		return i.(int) >= int(0)
	case "int32":
		return i.(int32) >= int32(0)
	case "int64":
		return i.(int64) >= int64(0)
	default:
		return i.(int) >= 0
	}
}

func validTaxCode(code int) bool {
	_, status := mapTaxCode(code)
	return status
}

func mapTaxCode(code int) (string, bool) {
	x, status := constant.TAXCODE_TYPE[code]
	return x, status
}
