package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

//Website struct information
type Website struct {
	Name string
	Time string
}

func main() {
	website := Website{"Mitch", time.Now().Format(time.Stamp)}
	templates := template.Must(template.ParseFiles("templates/website-template.html"))
	http.Handle("/static",
		http.StripPrefix("/static",
			http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if name := r.FormValue("name"); name != "" {
			website.Name = name
		}

		if err := templates.ExecuteTemplate(w, "website-template.html", website); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
