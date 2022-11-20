package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	wd += "\\cmd\\web\\templates\\"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, wd, "test.page.gohtml")
	})

	fmt.Println("Starting front end service on port 80")
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		log.Panic(err)
	}
}

func render(w http.ResponseWriter, wd, t string) {

	partials := []string{
		"base.layout.gohtml",
		"header.partial.gohtml",
		"footer.partial.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("%s/%s", wd, t))

	for _, x := range partials {
		templateSlice = append(templateSlice, fmt.Sprintf("%s/%s", wd, x))
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
