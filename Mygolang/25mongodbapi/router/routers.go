package router

import (
	"mongodbapi/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")

	return router
}
