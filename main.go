package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)


type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json") // this is to set the header of the response to json
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)  // this is to get the params from a request; it holds it in some sort of map, so you access id like params["id"]
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)  // this is the way you delete from a slice in Go; you basically join slices of the original slice
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}



func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "428231", Title: "Movie one", Director: &Director{Firstname: "Cactus", Lastname: "Jack"}}) // creating a bunch of dummy movies
	movies = append(movies, Movie{ID: "2", Isbn: "430452", Title: "Movie two", Director: &Director{Firstname: "Chibuoyim", Lastname: "Onuigwe"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))  // this logs an error with starting up the server, if there is one
}