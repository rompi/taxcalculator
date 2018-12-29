package logic

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/rompi/tax-calc/app/model"
)

type DB struct {
	*sql.DB
}

func (db DB) CreateObject(o *model.Object) (*model.Object, error) {
	stmt := fmt.Sprintf("INSERT INTO tax_object(name, tax_code, price) VALUES('%s',%d,%0.2f) RETURNING ID", o.Name, o.TaxCode, o.Price)

	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	var id int
	err := db.QueryRow(stmt).Scan(&id)
	if err != nil {
		log.Println("error in execute insert query", stmt, err.Error())
		return nil, err
	} else {
		o.Id = int64(id)
	}
	return o, nil
}

func (db DB) GetObjects() ([]*model.Object, int, error) {
	var objects []*model.Object
	stmt := fmt.Sprintf("SELECT id, name, tax_code, price from tax_object order by id desc")

	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	rows, err := db.Query(stmt)
	if err != nil {
		log.Println("error in execute get query ", stmt, err.Error())
	}

	i := 0
	if rows != nil {
		for rows.Next() {
			var o model.Object
			err = rows.Scan(&o.Id, &o.Name, &o.TaxCode, &o.Price)
			if err != nil {
				log.Println("error in scan query result", err.Error())
				return objects, i, err
			}
			objects = append(objects, &o)
			i++
		}
	}
	return objects, i, nil
}
