package main

import (
	"fmt"
	"net/http"
)

func main() {
	// configure server
	server := &http.Server{Addr: "localhost:8080"}

	// Routing
	http.Handle("/", Auth(http.HandlerFunc(getMovieAuth))) // get Books dari Book-API

	// Run server
	fmt.Println("Server is running on localhost:8080")
	server.ListenAndServe()
}

// Middleware
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		var username, password, ok = r.BasicAuth()

		if r.Method == "POST" {
			if !ok {
				// Karena pertama nge akses dia pasti !ok jadi ini text pertama yang ditampilin di Body.
				w.Write([]byte("FUCK YOU. YOU'RE NOT SUPPOSED TO BE IN HERE"))
				w.WriteHeader(http.StatusOK)
				return
			}

			if username == "admin" && password == "admin" {
				w.WriteHeader(http.StatusOK)
				next.ServeHTTP(w, r)
				return
			}

			w.Write([]byte("Username atau Password tidak sesuai."))
			return
		}
	})

}

// Untuk menampilkan text di browser
func getMovieAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Write([]byte("<h1>Anda Berhasil Mengakses Fungsi GetMovie() </h1>"))
		
	}
}