package logic

import (
	"testing"

	"github.com/rompi/tax-calc/app/model"
	. "github.com/smartystreets/goconvey/convey"
)

func TestValidate(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given name, type, and taxcode is fine", t, func() {
		data := model.BuildObject("Barongsai", 1, float64(10000))

		Convey("When the data is validated", func() {
			res := Validate(data)
			Convey("The error message should be nil", func() {
				So(res, ShouldEqual, nil)
			})
		})
	})

	Convey("Given name is empty", t, func() {
		data := model.BuildObject("", 1, float64(10000))

		Convey("When the data is validated", func() {
			res := Validate(data)
			Convey("The error message should be exists", func() {
				So(res[0], ShouldEqual, "Name should not be empty")
			})
		})
	})

	Convey("Given tax code is invalid", t, func() {
		data := model.BuildObject("Barongsai", 6, float64(10000))

		Convey("When the data is validated", func() {
			res := Validate(data)
			Convey("The error message should be exists", func() {
				So(res[0], ShouldEqual, "Invalid Tax Code")
			})
		})
	})

	Convey("Given price is less than zero", t, func() {
		data := model.BuildObject("Barongsai", 6, float64(-1))

		Convey("When the data is validated", func() {
			res := Validate(data)
			Convey("The error message should be exists", func() {
				So(res[0], ShouldEqual, "Price must be greater than zero")
			})
		})
	})

	Convey("Given name, tax code, and price is not valid", t, func() {
		data := model.BuildObject("", 6, float64(-1))

		Convey("When the data is validated", func() {
			res := Validate(data)
			Convey("The error message should be exists 3 message", func() {
				So(len(res), ShouldEqual, 3)
			})
		})
	})

}
