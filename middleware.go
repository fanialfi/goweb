package main

import "net/http"

const (
	USERNAME = "fani"
	PASSWORD = "1234"
)

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte("terjadi kesalahan"))
		return false
	}

	isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		w.Write([]byte("wrong username or password"))
		return false
	}

	return true
}

func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodGet {
		w.Write([]byte("hanya mengizinkan GET"))
		return false
	}

	return true
}
