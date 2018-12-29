package svc

import "github.com/rompi/tax-calc/app/model"

type Bill interface {
	Create(*model.Object) (*model.Object, error)
	Read() ([]*model.Object, int, error)
}
