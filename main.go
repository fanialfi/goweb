package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", routeIndexGet)
	http.HandleFunc("/process", routeSubmitPost)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func routeIndexGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.New("form").ParseFiles(filepath.Join("views", "view.html")))
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func routeSubmitPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header.Get("Content-Type"))
	if r.Method == http.MethodPost {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		tmpl := template.Must(template.New("result").ParseFiles(filepath.Join("views", "view.html")))

		if err := r.ParseMultipartForm(10 << 10); err != nil {
			fmt.Println("error cok", err.Error())
			http.Error(w, fmt.Sprintf("error parse form : %s", err.Error()), http.StatusInternalServerError)
			return
		}

		name := r.FormValue("name")
		message := r.FormValue("message")
		data := map[string]string{
			"name":    name,
			"message": message,
		}

		if err := tmpl.Execute(w, data); err != nil {
			fmt.Println("error cok", err.Error())
			http.Error(w, fmt.Sprintf("anjay error : %s", err.Error()), http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}
