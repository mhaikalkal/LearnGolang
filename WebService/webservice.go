package main

import (
	"encoding/json"
	"fmt" // mau coba bikin API Dulu, nanti di enable lagi
	"net/http"
)

func main() {
    // pattern itu untuk URL nya. baseURL lalu patern. contoh :::
    http.HandleFunc("/users", users) // http://localhost:8080/users. Diisi dengan func users
    http.HandleFunc("/user", user) // http://localhost:8080/use. Diisi dengan func user

    fmt.Println("starting web server at http://localhost:8080/") // cuma penanda doang
    http.ListenAndServe(":8080", nil) // deploy
}

type student struct {
    ID    string
    Name  string
    Grade int
}

// sample data
var data = []student{
    student{"E001", "ethan", 21},
    student{"W001", "wick", 22},
    student{"B001", "bourne", 23},
    student{"B002", "bond", 23},
}

func users(w http.ResponseWriter, r *http.Request) { // seluruh isi struct []student
	// Statement w.Header().Set("Content-Type", "application/json") digunakan untuk menentukan tipe response. Yaitu sbg JSON
	w.Header().Set("Content Type", "application/json")

	// Jika pengguna meminta data "GET". maka:
	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		// Kalau error maka tunjukkan: response, pesan error, error code
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return // berhenti/ keluar dari if.
		}

		// Jika tidak error
		// Sedangkan w.Write() digunakan untuk mendaftarkan data sebagai response.
		w.Write(result) // response-nya result []byte struct data
		return // berhenti
	}

	// jika proses tidak valid.
	http.Error(w, "", http.StatusBadRequest)
}

func user(w http.ResponseWriter, r *http.Request) { // 1 data struct []student
    w.Header().Set("Content-Type", "application/json") // response sebagai json

	// jika proses "GET". atau minta data, maka:
    if r.Method == "GET" {
        var id = r.FormValue("id") // penampung. show data struct []student berdasarkan id
        var result []byte // hasilnya sebagai []byte. untuk di jadikan json.Marshal()
        var err error // penampung error

		// karena data merupakan slice, bisa di loop dulu
        for _, each := range data {
			// Jika setiap field ID == var id. 
			// bahasa manusia: jika id yg dimasukkan cocok dengan value ID, maka:
            if each.ID == id {
				// masukkan data ke result, dan error ke var err
                result, err = json.Marshal(each)

				// Jika gagal maka:
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }

				// Jika sukses maka
                w.Write(result) // response-nya berupa json
                return
            }
        }

		// jika data tidak ada. throw:
		// response, pesan error, status errorNotFound
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

	// Jika ada kesalahan.
    http.Error(w, "", http.StatusBadRequest)
}