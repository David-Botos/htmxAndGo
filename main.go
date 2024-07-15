package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("hello world. The time is", time.Now())

	// w = write and r = read
	h1 := func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("index.html"))

		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "Meme", Director: "Schmeme"},
			},
		}

		template.Execute(w, films) // nil would be the place you place context like django
	}
	http.HandleFunc("/", h1)

	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		htmlStr := fmt.Sprintf("<li><p>{{ .Title }} %s - %s {{ .Director }}</p></li>", title, director)
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, nil)
	}
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
