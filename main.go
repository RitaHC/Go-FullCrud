package main

import (
	"fmt"
	"log"

	// for parsing data to and from json
	"encoding/json"
	// to create a new id for the Create in Crud

	"net/http"

	// for converting new id to string

	"github.com/gorilla/mux"
)

type Movie struct {
	Released string `json: "Release Date"`
	Title    string `json: "Title"`
	Genre    string `json: "Genre"`
	Country  string `json: "Country"`
	ImdbID   string `json: "imdbID"`
	Director string `json: "Director"`
	Language string `json: "Language"`
}

var movies []Movie

// var movies []Movie

func apiCalling() ([]Movie, error) {
	// Step 1 : Call api
	apiURL := "https://gist.githubusercontent.com/saniyusuf/406b843afdfb9c6a86e25753fe2761f4/raw/523c324c7fcc36efab8224f9ebb7556c09b69a14/Film.JSON"
	response, err := http.Get(apiURL)
	// error handling
	if err != nil {
		return nil, err
	}
	// Step 2 : Decode the data coming form API
	if err := json.NewDecoder(response.Body).Decode(&movies); err != nil {
		return nil, err
	}
	return movies, nil

}

//====================== Get All ============================

func getMovies(w http.ResponseWriter, r *http.Request) {
	//Step 1 : Fetch all movies from the API whcih are decoded
	movies, err := apiCalling()
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
	}
	// Step 2 : Send data to display
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

// ================================== Delete One ==================================
func deleteMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range movies {
		// now find the particular movie with params
		if movie.ImdbID == params["id"] {
			fmt.Println("DELETED : ", movie)
			// delete that id
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// ======================== Get One ========================
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, movie := range movies {
		if movie.ImdbID == params["id"] {
			fmt.Println("One Movie Found", movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

//======================== Create Movie ===================

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Create a new movie instance to be send as body to postman
	var movie Movie
	// Decode the new movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movies = append(movies, movie)
	fmt.Println("New Movie added", movie)
	// return the new movie created
	json.NewEncoder(w).Encode(movie)
}

// ====================== Update Movie ===================
func updateMovie(w http.ResponseWriter, r *http.Request) {

	// set json content type
	w.Header().Set("Content-Type", "application/json")
	// params
	params := mux.Vars(r)
	// loop over movies

	// delete the movie with the id that we have send
	// add a new movie - which is the movie we've send via postman
	for index, item := range movies {
		//Step 1 : find and delete the old movie
		if item.ImdbID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			// Step 2:  Now create new movie
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ImdbID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}

	}
}

func main() {
	// create a new router
	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting Server at PORT 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
