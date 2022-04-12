package main

import (
	"fmt"
	"net/http"
)

/*
	Middleware adalah satu satu mekanisme keamanan website. Middleware berjalan sebelum fungsi utama di proses.
	Sebagai contoh serderhana, misalnya terdapat routing untuk mengarahkan ke dalam fungsi, fungsi tersebut digunakan untuk melihat data. Nah agar fungsi tersebut tidak dapat di akses orang maka harus melewati middleware, middleware berperan untuk menghentikan atau meneruskan proses.
	Contoh paling simpel isi dari middleware adalah pemeriksaan suatu teks yang bisa di kirim melalui body dan header melalui http request.
	Untuk the real word projek middleware di gunakan untuk auntentifikasi login menggunakan token, middleware berperan untuk memeriksa apakah token yang dikirim sesuai atau tidak.
*/

func main() {
	// konfigurasi server
	server := &http.Server{ Addr: "localhost:8080" }
  
	// routing
	http.Handle("/", Log(http.HandlerFunc(GetMovie)))
  
	// jalankan server
	fmt.Println("server running at http://localhost:8080")
	server.ListenAndServe()
  }

type Movie struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
  }
  
func Movies() []Movie {
	movs := []Movie{
		{ 1, "Spider-Man", 2002 },
		{ 2, "John Wick", 2014 },
		{ 3, "Avengers", 2018 },
		{ 4, "Logan", 2017 },
	}

	return movs
}

// fungsi log berfungsi sebagai middleware
func Log(next http.Handler) http.Handler{
	return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Ini dari middleware Log....\n")
		fmt.Println(r.URL.Path)

		next.ServeHTTP(w, r)
	})
}

// Fungsi GetMovie untuk mengampilkan text string di browser
func GetMovie(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<h1>Anda Berhasil Mengakses Fungsi GetMovie() </h1>"))
}