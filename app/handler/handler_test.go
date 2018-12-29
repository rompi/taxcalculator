package handler

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rompi/tax-calc/app/model"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHandleCreateObject(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Test Normal Handle Create Object Type 1", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mock := NewMockBill(mockCtrl)

		l := &Handler{
			Log:   logrus.New(),
			Logic: mock,
		}
		data := model.BuildObject("TEST", 1, 100)
		mock.EXPECT().Create(data).Return(data, nil)
		Convey("When the data is validated", func() {
			statusCode, res := HandleCreateObject(*l, data)
			Convey("The error message should not be exists", func() {
				So(res.Data, ShouldNotBeNil)
				So(statusCode, ShouldEqual, 201)
			})
		})
	})

	// Only pass t into top-level Convey calls
	Convey("Test Failed Empty Name Handle Create Object", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mock := NewMockBill(mockCtrl)

		l := &Handler{
			Log:   logrus.New(),
			Logic: mock,
		}
		data := model.BuildObject("", 2, 100)

		Convey("When the data is validated", func() {
			statusCode, res := HandleCreateObject(*l, data)
			Convey("The error message should be exists", func() {
				So(res.Error[0], ShouldEqual, "Name should not be empty")
				So(statusCode, ShouldEqual, 400)
			})
		})
	})

	// Only pass t into top-level Convey calls
	Convey("Test Failed Invalid TaxCode Handle Create Object", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mock := NewMockBill(mockCtrl)

		l := &Handler{
			Log:   logrus.New(),
			Logic: mock,
		}
		data := model.BuildObject("TEST", 9999, 100)

		Convey("When the data is validated", func() {
			statusCode, res := HandleCreateObject(*l, data)
			Convey("The error message should be exists", func() {
				So(res.Error[0], ShouldEqual, "Invalid Tax Code")
				So(statusCode, ShouldEqual, 400)
			})
		})
	})

	// Only pass t into top-level Convey calls
	Convey("Test Failed Negative Price Create Object", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mock := NewMockBill(mockCtrl)

		l := &Handler{
			Log:   logrus.New(),
			Logic: mock,
		}
		data := model.BuildObject("TEST", 1, -100)

		Convey("When the data is validated", func() {
			statusCode, res := HandleCreateObject(*l, data)
			Convey("The error message should be exists", func() {
				So(res.Error[0], ShouldEqual, "Price must be greater than zero")
				So(statusCode, ShouldEqual, 400)
			})
		})
	})
	// Only pass t into top-level Convey calls
	Convey("Test Failed When Create Object to database", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mock := NewMockBill(mockCtrl)

		l := &Handler{
			Log:   logrus.New(),
			Logic: mock,
		}
		data := model.BuildObject("TEST", 1, 100)
		mock.EXPECT().Create(data).Return(nil, fmt.Errorf("gagal insert"))

		Convey("When the data is validated", func() {
			statusCode, res := HandleCreateObject(*l, data)
			Convey("The error message should be exists", func() {
				So(res.Error[0], ShouldEqual, "gagal insert")
				So(statusCode, ShouldEqual, 500)
			})
		})
	})
}

func TestHandleGetObjects(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Test Normal Handle Get Object", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mock := NewMockBill(mockCtrl)

		l := &Handler{
			Log:   logrus.New(),
			Logic: mock,
		}
		data := model.BuildObjects(model.BuildObject("TEST", 1, 1000))
		mock.EXPECT().Read().Return(data, 1, nil)
		Convey("When the data is validated", func() {
			statusCode, res := HandleGetObject(*l)
			Convey("The error message should not be exists", func() {
				So(res.Data, ShouldNotBeNil)
				So(statusCode, ShouldEqual, 200)
			})
		})
	})

	// Only pass t into top-level Convey calls
	Convey("Test Failed Handle Get Object", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mock := NewMockBill(mockCtrl)

		l := &Handler{
			Log:   logrus.New(),
			Logic: mock,
		}
		data := model.BuildObjects(model.BuildObject("TEST", 1, 1000))
		mock.EXPECT().Read().Return(data, 1, fmt.Errorf("gagal get"))
		Convey("When the data is validated", func() {
			statusCode, res := HandleGetObject(*l)
			Convey("The error message should be exists", func() {
				So(res.Error, ShouldNotBeNil)
				So(statusCode, ShouldEqual, 500)
			})
		})
	})
}
