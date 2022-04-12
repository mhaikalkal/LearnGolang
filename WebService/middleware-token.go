package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Konfigurasi Server
	server := &http.Server{ Addr: "localhost:8080", }

	// Routing
	http.Handle("/", CekToken(http.HandlerFunc(getMovie)))

	// Jalankan Server
	fmt.Println("server running at http://localhost:8080")
	server.ListenAndServe()


}

func CekToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		
		if r.URL.Query().Get("token") != "42069" {
			// Jika kode token tidak sesuai. maka return (stop)
			fmt.Fprintf(w, "Token tidak sesuai\n")
			return
		}

		// jika sesuai maka lanjut
		next.ServeHTTP(w, r)
	})
}


func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HAHAHAH")) // test duls
}

