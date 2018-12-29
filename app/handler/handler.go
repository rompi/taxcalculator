package handler

import (
	"net/http"

	"github.com/rompi/tax-calc/app/logic"
	"github.com/rompi/tax-calc/app/model"
	"github.com/rompi/tax-calc/app/svc"
	"github.com/sirupsen/logrus"
)

// Handler logic and log inside
type Handler struct {
	Logic svc.Bill
	Log   *logrus.Logger
}

// HandleCreateObject and build response message
func HandleCreateObject(h Handler, o *model.Object) (int, *model.Response) {
	// validate input value
	errVal := validateInput(o)
	if len(errVal) > 0 {
		payload := model.BuildResponse(nil, errVal)
		return http.StatusBadRequest, payload
	}

	// processing create new data
	data, err := h.Logic.Create(o)
	if err != nil {
		payload := model.BuildResponse(o, []string{err.Error()})
		return http.StatusInternalServerError, payload
	}
	payload := model.BuildResponse(data, nil)
	return http.StatusCreated, payload
}

// HandleGetObject and build response message
func HandleGetObject(h Handler) (int, *model.ResponseList) {
	data, _, err := h.Logic.Read()
	if err != nil {
		payload := model.BuildResponseList(data, []string{err.Error()})
		return http.StatusInternalServerError, payload
	}

	payload := model.BuildResponseList(data, nil)

	return http.StatusOK, payload
}

func validateInput(o *model.Object) []string {
	return logic.Validate(o)
}
