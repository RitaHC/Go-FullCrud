package main

import (
	"fmt"
	"log"

	// for parsing data to and from json
	"encoding/json"
	// to create a new id for the Create in Crud
	"math/rand"
	"net/http"

	// for converting new id to string
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json: "id"`
	Isbn     string    `json: "isbn"`
	Title    string    `json: "title"`
	Director *Director `json: "director"`
}

type Director struct {
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
}

var movies []Movie

// Get all
func getMovies(w http.ResponseWriter, r *http.Request) {
	// set content type as json
	// then our code converts json to its struct
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// delete function
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		// now find the particular movie with params
		if item.ID == params["id"] {
			fmt.Println(item)
			// delete that id
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

// Get One Movie
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		fmt.Println(item)
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// Create Movie

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Create a new movie instance to be send as body to postman
	var movie Movie
	// Decod ethe new movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// create a id for the movie and convert it to string to be used later
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	fmt.Println("New Movie added", movie)
	// return the new movie created
	json.NewEncoder(w).Encode(movie)
}

// Update Movie
func updateMovie() {

	// set json content type
	w.Header().Set("Content-Type", "application/json")
	// params
	params := mux.Vars(r)
	// loop over movies

	// delete the movie with the id that we have send
	// add a new movie - which is the movie we've send via postman

}

func main() {
	// create a new router
	r := mux.NewRouter()

	// Seed data
	movies = append(movies, Movie{ID: "1", Isbn: "376", Title: "Movie One", Director: &Director{Firstname: "Hameer", Lastname: "Chauhan"}})
	movies = append(movies, Movie{ID: "2", Isbn: "264", Title: "Movie Two", Director: &Director{Firstname: "Rajat", Lastname: "Kumar"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/moies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting Server at PORT 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
