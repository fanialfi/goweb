package main

// middleware menggunakan DefaultServeMux

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	// http.ServeMux adalah sebuah router yang semua router dibuat akan didaftarkan di sini
	mux := http.DefaultServeMux

	mux.HandleFunc("/student", ActionStudent)

	// registrasi middleware
	// karena http.ServeMux juga mengimplementasikan interface http.Handler
	// http.serveMux digunakan untuk menghandle semua permintaan dari clien dan diteruskan ke router terkait
	var handler http.Handler = mux
	handler = MiddlewareAuth(handler)
	handler = MiddlewareAllowOnlyGET(handler)

	// http.Server adalah server http-nya itu sendiri yang bertanggung jawab
	// untuk mendengarkan request dari port tertentu dan kemudian request itu akan diteruskan ke Handler
	// dalam hal ini request akan diteruskan ke ServeMux
	server := new(http.Server)
	server.Addr = ":9000"
	server.Handler = handler

	log.Println("server started at localhost:9000")
	server.ListenAndServe()
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, GetStudents())
}

func OutputJSON(w http.ResponseWriter, o any) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
