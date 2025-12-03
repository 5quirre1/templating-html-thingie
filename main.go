package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("greg.html"))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Name string
		}{
			Name : "greg"
		}
		if err := tmpl.Execute(w, data); err != nil {
			log.Println(err)
		}
	})
	http.ListenAndServe(":808-", nil)
}
