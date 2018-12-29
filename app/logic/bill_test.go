package logic

import (
	"errors"
	"log"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rompi/tax-calc/app/model"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBillLogicCreateObject(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Test Normal Create Object Type 1", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockDB := NewMockDBTrx(mockCtrl)

		bill := &Billing{
			Log:      logrus.New(),
			Database: mockDB,
		}
		data := model.BuildObject("TEST", 1, 100)

		mockDB.EXPECT().CreateObject(data).Return(data, nil)
		Convey("When the data is validated", func() {
			res, err := bill.Create(data)
			Convey("The error message should not be exists", func() {
				So(err, ShouldBeNil)
				So(res.Tax, ShouldEqual, float64(10))
				So(res.Amount, ShouldEqual, float64(110))
			})
		})
	})

	// Only pass t into top-level Convey calls
	Convey("Test Normal Create Object Type 2", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockDB := NewMockDBTrx(mockCtrl)

		bill := &Billing{
			Log:      logrus.New(),
			Database: mockDB,
		}
		data := model.BuildObject("TEST", 2, 100)

		mockDB.EXPECT().CreateObject(data).Return(data, nil)
		Convey("When the data is validated", func() {
			res, err := bill.Create(data)
			Convey("The error message should not be exists", func() {
				So(err, ShouldBeNil)
				So(res.Tax, ShouldEqual, float64(12))
				So(res.Amount, ShouldEqual, float64(112))
			})
		})
	})

	// Only pass t into top-level Convey calls
	Convey("Test Normal Create Object Type 3", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockDB := NewMockDBTrx(mockCtrl)

		bill := &Billing{
			Log:      logrus.New(),
			Database: mockDB,
		}
		data := model.BuildObject("TEST", 3, 90)

		mockDB.EXPECT().CreateObject(data).Return(data, nil)
		Convey("When the data is validated", func() {
			res, err := bill.Create(data)
			Convey("The error message should not be exists", func() {
				So(err, ShouldBeNil)
				So(res.Tax, ShouldEqual, float64(0))
				So(res.Amount, ShouldEqual, float64(90))
			})
		})
	})

	// Only pass t into top-level Convey calls
	Convey("Test Normal Create Object Type 3 (cont.)", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockDB := NewMockDBTrx(mockCtrl)

		bill := &Billing{
			Log:      logrus.New(),
			Database: mockDB,
		}
		data := model.BuildObject("TEST", 3, 1000)

		mockDB.EXPECT().CreateObject(data).Return(data, nil)
		Convey("When the data is validated", func() {
			res, err := bill.Create(data)
			Convey("The error message should not be exists", func() {
				So(err, ShouldBeNil)
				So(res.Tax, ShouldEqual, float64(9))
				So(res.Amount, ShouldEqual, float64(1009))
			})
		})
	})

	// Only pass t into top-level Convey calls
	Convey("Test Failed Create Object", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockDB := NewMockDBTrx(mockCtrl)

		bill := &Billing{
			Log:      logrus.New(),
			Database: mockDB,
		}
		data := model.BuildObject("TEST", 1, 1)

		mockDB.EXPECT().CreateObject(data).Return(data, errors.New("gagal insert"))
		Convey("When the data is validated", func() {
			res, err := bill.Create(data)
			Convey("The error message should be exists", func() {
				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})
		})
	})

	// Only pass t into top-level Convey calls
	Convey("Test Failed Create Object", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockDB := NewMockDBTrx(mockCtrl)

		bill := &Billing{
			Log:      logrus.New(),
			Database: mockDB,
		}
		data := model.BuildObject("TEST", 1, 1)

		mockDB.EXPECT().CreateObject(data).Return(data, errors.New("gagal insert"))
		Convey("When the data is validated", func() {
			res, err := bill.Create(data)
			log.Println(res, err)
			Convey("The error message should be exists", func() {
				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})
		})
	})

}

func TestBillLogicReadObject(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Test Normal Get Object Type 1", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockDB := NewMockDBTrx(mockCtrl)

		bill := &Billing{
			Log:      logrus.New(),
			Database: mockDB,
		}
		data := model.BuildObjects(model.BuildObject("TEST", 1, 100))

		mockDB.EXPECT().GetObjects().Return(data, 1, nil)
		Convey("When the data is validated", func() {
			res, i, err := bill.Read()
			Convey("The error message should not be exists", func() {
				So(err, ShouldBeNil)
				So(res[0].Tax, ShouldEqual, float64(10))
				So(res[0].Amount, ShouldEqual, float64(110))
				So(i, ShouldEqual, 1)
			})
		})
	})

	// Only pass t into top-level Convey calls
	Convey("Test Normal Get Object Type 2", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockDB := NewMockDBTrx(mockCtrl)

		bill := &Billing{
			Log:      logrus.New(),
			Database: mockDB,
		}
		data := model.BuildObjects(model.BuildObject("TEST", 2, 100))

		mockDB.EXPECT().GetObjects().Return(data, 1, nil)
		Convey("When the data is validated", func() {
			res, _, err := bill.Read()
			Convey("The error message should not be exists", func() {
				So(err, ShouldBeNil)
				So(res[0].Tax, ShouldEqual, float64(12))
				So(res[0].Amount, ShouldEqual, float64(112))
			})
		})
	})

	// Only pass t into top-level Convey calls
	Convey("Test Normal Get Object Type 3", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockDB := NewMockDBTrx(mockCtrl)

		bill := &Billing{
			Log:      logrus.New(),
			Database: mockDB,
		}
		data := model.BuildObjects(model.BuildObject("TEST", 3, 90))

		mockDB.EXPECT().GetObjects().Return(data, 1, nil)
		Convey("When the data is validated", func() {
			res, _, err := bill.Read()
			Convey("The error message should not be exists", func() {
				So(err, ShouldBeNil)
				So(res[0].Tax, ShouldEqual, float64(0))
				So(res[0].Amount, ShouldEqual, float64(90))
			})
		})
	})

	// Only pass t into top-level Convey calls
	Convey("Test Data is null", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockDB := NewMockDBTrx(mockCtrl)

		bill := &Billing{
			Log:      logrus.New(),
			Database: mockDB,
		}

		mockDB.EXPECT().GetObjects().Return(nil, 0, nil)
		Convey("When the data is validated", func() {
			res, _, err := bill.Read()
			Convey("The error message should not be exists", func() {
				So(err, ShouldBeNil)
				So(res, ShouldBeNil)
			})
		})
	})

}
