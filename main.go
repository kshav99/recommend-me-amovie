package main

import (
	"fmt"
	"net/http"
	"recommend-me-amovie/api"
)

func main() {
	http.HandleFunc("/api/genres", api.Genres)
	http.HandleFunc("/api/movie-types", api.MovieTypes)
	http.HandleFunc("/api/countries", api.Countries)
	http.HandleFunc("/api/recommend", api.Recommend)
	http.HandleFunc("/api/user-country", api.UserCountry)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "static/index.html")
	})
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}
