package main

import(
	"fmt"
	"log"
	"encoding/json"
	// "os"
	// "math/rand"
	"net/http"
	// "strconv"
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

func getMovies( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	json.NewEncoder(w).Encode(&Movie{})

}

func createMovie( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for i, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			var movie Movie
			json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for i, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}


func main(){
	r:=mux.NewRouter()

	movies=append(movies, Movie{ID:"1",Isbn: "1a",Title: "Movie 1", Director: &Director{Firstname: "Ayush", Lastname: "Gupta"}})
	movies=append(movies, Movie{ID:"2",Isbn: "2a",Title: "Movie 2", Director: &Director{Firstname: "Abhinav", Lastname: "Gupta"}})

	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movie/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at PORT 8080")

	log.Fatal(http.ListenAndServe(":8080",r))

}

