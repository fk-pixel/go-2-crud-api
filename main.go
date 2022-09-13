package main

import (
	"encoding/json"
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
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

func createMovie(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "application/json")
var movie Movie 
_ = json.NewDecoder(r.body).Decode(&movie)
movie.ID = strconv.Itoa(rand.Intn(10000000))
movies = append(movies, movie)
json.NewEncoder(w).Encode(movie)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := max.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies) 
}

func updateMovie(w http.Response, r *http.Request){
	w.Header().Set()
}


func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{{ID: "1", Isbn: "123456789", Title: "First Movie", Director: &Director{Firstname: "Adam", Lastname: "Markdown"}})
	movies = append(movies, Movie{{ID: "2", Isbn: "012345678", Title: "Second Movie", Director: &Director{Firstname: "Kurt", Lastname: "Agent"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/[id]", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/[id]", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/[id]", updateMovie).Methods("PUT")

	fmt.Print("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
