package main

import (
	"encoding/json"
	"fmt"
	"log"
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

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
func createMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies { //for each
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "500", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "501", Title: "Movie Two", Director: &Director{Firstname: "Halim", Lastname: "Ben Oun"}})
	r.HandleFunc("/movies", getMovies).Methods("get")
	r.HandleFunc("/movies/{id}", getMovie).Methods("get")
	r.HandleFunc("/movies", createMovies).Methods("post")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("put")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("delete")
	fmt.Printf("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// func formHandler(w http.ResponseWriter, r *http.Request) {
// 	if err := r.ParseForm(); err != nil {
// 		fmt.Fprintf(w, "parse form", err)
// 		return
// 	}
// 	fmt.Fprintf(w, "post successful")
// 	name := r.FormValue("name")
// 	adress := r.FormValue("adress")
// 	fmt.Fprintf(w, "name - %s", name)
// 	fmt.Fprintf(w, "adress: %s", adress)
// }
// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/hello" {
// 		http.Error(w, "404 not found", http.StatusNotFound)
// 	}
// 	if r.Method != "GET" {
// 		http.Error(w, "method not supported", http.StatusNotFound)
// 	}
// 	fmt.Fprintf(w, "hello world")
// }
// func main() {
// 	fileServer := http.FileServer(http.Dir("./static"))
// 	http.Handle("/", fileServer)
// 	http.HandleFunc("/form", formHandler)
// 	http.HandleFunc("/hello", helloHandler)
// 	fmt.Printf("Starting server at port 8080 ")
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal(err)
// 	}
// }
