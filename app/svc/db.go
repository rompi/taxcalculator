package svc

import "github.com/rompi/tax-calc/app/model"

type DBTrx interface {
	CreateObject(*model.Object) (*model.Object, error)
	GetObjects() ([]*model.Object, int, error)
}
