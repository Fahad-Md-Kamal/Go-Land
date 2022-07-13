package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"math/rand"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json: "id"`
	Isbn string `json: "isbn"`
	Title string `json: "title"`
	Director *Director `json: "director"`
}

type Director struct {
	Firstname string `json: "firstname"`
	Lastname string `json: "lastname"`
}

var movies []Movie


func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}


func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, items := range movies{

		if items.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}


func getMovieDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, items := range movies{
		if items.ID == params["id"]{
			json.NewEncoder(w).Encode(items)
			return
		}
	}
}


func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}


func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var newMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&newMovie)
	newMovie.ID = params["id"]
	for index, items := range movies{
		if items.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			movies = append(movies, newMovie)
			json.NewEncoder(w).Encode(movies)
		}
	}
}


func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "4321", Title: "Star Dust", Director: &Director{Firstname: "James", Lastname: "Petersone"}})
	movies = append(movies, Movie{ID: "2", Isbn: "98765", Title: "Marvel World", Director: &Director{Firstname: "James", Lastname: "Cameron"}})

	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovieDetail).Methods("GET")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Server is listening at port:8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}