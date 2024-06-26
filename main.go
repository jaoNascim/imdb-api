package main

import (
	"html/template"
	"net/http"
	"strconv"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("pages/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/process", processor)
	http.ListenAndServe(":8080", nil)

	getMovieData("Whiplash")
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func processor(w http.ResponseWriter, r *http.Request) {
	movie := r.FormValue("movieName")

	moviePoster := getMoviePoster(movie)
	movieName := getMovieName(movie)
	movieYear := strconv.Itoa(getMovieYear(movie))

	d := struct {
		MovieName   string
		MoviePoster string
		MovieYear   string
	}{
		MovieName:   movieName,
		MoviePoster: moviePoster,
		MovieYear:   movieYear,
	}

	tpl.ExecuteTemplate(w, "processor.gohtml", d)
}
