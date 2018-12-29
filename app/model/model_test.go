package model

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestModel(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Test Build Response OK", t, func() {
		o := BuildObject("TEST", 1, 1000)
		Convey("When the data is validated", func() {
			res := BuildResponse(o, nil)
			Convey("The error message should be exists", func() {
				So(res.Data.Name, ShouldEqual, "TEST")
				So(res.Error, ShouldBeNil)
			})
		})
	})

	// Only pass t into top-level Convey calls
	Convey("Test Build ResponseList OK", t, func() {
		o := BuildObject("TEST", 1, 1000)
		ol := BuildObjects(o)
		Convey("When the data is validated", func() {
			res := BuildResponseList(ol, nil)
			Convey("The error message should be exists", func() {
				So(res.Data[0].Name, ShouldEqual, "TEST")
				So(res.Error, ShouldBeNil)
			})
		})
	})
	// Only pass t into top-level Convey calls
	Convey("Test Build Object OK", t, func() {
		Convey("When the data is validated", func() {
			res := BuildObject("TEST", 1, 1000)
			Convey("The error message should be exists", func() {
				So(res.Name, ShouldEqual, "TEST")
			})
		})
	})

	// Only pass t into top-level Convey calls
	Convey("Test Build ResponseList OK", t, func() {
		Convey("When the data is validated", func() {
			o := BuildObject("TEST", 1, 1000)
			res := BuildObjects(o)
			Convey("The error message should be exists", func() {
				So(res[0].Name, ShouldEqual, "TEST")
			})
		})
	})
}
