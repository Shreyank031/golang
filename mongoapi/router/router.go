package router

import (
	"github.com/Shreyank031/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("api/movies", controller.GetAllMovies).Methods("GET")        //Get all movie
	router.HandleFunc("api/movie", controller.CreateMovie).Methods("POST")         //Creating a movie
	router.HandleFunc("api/movie/{id}", controller.MarkedAsWatched).Methods("PUT") //update watched
	router.HandleFunc("api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("api/deleteallmovie", controller.GetAllMovies).Methods("DELETE")

	return router
}
