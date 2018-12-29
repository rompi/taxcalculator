package logic

import (
	"log"

	"github.com/rompi/tax-calc/app/model"
	"github.com/rompi/tax-calc/app/svc"
	"github.com/rompi/tax-calc/constant"
	"github.com/sirupsen/logrus"
)

type Billing struct {
	Log      *logrus.Logger
	Database svc.DBTrx
}

func convertTax3(price float64) float64 {
	if 0 < price && price < 100 {
		return 0
	}
	return 0.01 * (price - 100)

}

func convertTax2(price float64) float64 {
	return 10 + (0.02 * price)
}

func convertTax1(price float64) float64 {
	return 0.1 * (price)
}

func calculateTax(o *model.Object) float64 {
	switch code := o.TaxCode; code {
	case 1:
		return convertTax1(o.Price)
	case 2:
		return convertTax2(o.Price)
	case 3:
		return convertTax3(o.Price)
	default:
		return 0
	}
}

func convertObject(o *model.Object) {
	o.Type = constant.TAXCODE_TYPE[o.TaxCode]
	o.Refundable = constant.IS_REFUNDABLE[o.TaxCode]
	o.Tax = calculateTax(o)
	o.Amount = o.Price + o.Tax
}

func (b Billing) Create(o *model.Object) (*model.Object, error) {

	o, err := b.Database.CreateObject(o)
	log.Println(err)
	if err != nil {
		b.Log.Errorf(err.Error())
		return nil, err
	}
	convertObject(o)

	return o, nil
}

func (b Billing) Read() ([]*model.Object, int, error) {
	objects, i, err := b.Database.GetObjects()
	for _, object := range objects {
		convertObject(object)
	}
	return objects, i, err
}
