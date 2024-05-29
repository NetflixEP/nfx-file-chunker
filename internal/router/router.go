package router

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/uploadContent", uploadContent).Methods("POST")
	return router
}
