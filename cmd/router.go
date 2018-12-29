package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rompi/tax-calc/app/handler"
	"github.com/rompi/tax-calc/app/model"
)

func routers(h handler.Handler) {
	router.Post("/object", func(w http.ResponseWriter, r *http.Request) {
		var o *model.Object
		err := json.NewDecoder(r.Body).Decode(&o)
		if err != nil {
			payload := model.BuildResponse(nil, []string{err.Error()})
			jsonResponse(w, http.StatusBadRequest, payload)
			return
		}

		statusCode, payload := setObject(h, o)
		jsonResponse(w, statusCode, payload)
		return
	})

	router.Get("/object", func(w http.ResponseWriter, r *http.Request) {
		statusCode, payload := getObject(h)
		jsonResponse(w, statusCode, payload)
		return
	})

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		payload := model.BuildResponse(nil, []string{"404 page not found"})
		statusCode := http.StatusNotFound
		jsonResponse(w, statusCode, payload)
		return
	})
}

func setObject(h handler.Handler, o *model.Object) (int, *model.Response) {
	statusCode, payload := handler.HandleCreateObject(h, o)
	return statusCode, payload

}

func getObject(h handler.Handler) (int, *model.ResponseList) {
	statusCode, payload := handler.HandleGetObject(h)
	return statusCode, payload
}

// jsonResponse write json response format
func jsonResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload, response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
