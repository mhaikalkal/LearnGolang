package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	// Routing
	http.HandleFunc("/books", GetBooks)
	http.HandleFunc("/book", GetBookById)
	http.HandleFunc("/post", AddBookViaPost)

	// Configure Server + Run Server
	fmt.Println("starting web server at http://localhost:8080/")
	// menerima *http.Server. Jadi bisa aja kita buat struct. Misal : var server = &http.Server{ Addr: "localhost:8080" }
	http.ListenAndServe("localhost:8080", nil)

}

type Books struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Desc        string `json:"desc"`
	Author      string `json:"author"`
	ReleaseYear int    `json:"release-year"`
}

func BookDatas() []Books {
	var BooksData = []Books{
		{Id: 1, Title: "Overlord", Desc: "n 21st century, world entered a new stage of VR games…and “YGGDRASIL” is considered top of all MMORPG…but, After announcing that all its servers will be off, the internet game Yggdrasil shut down…or so was supposed to happen, but for some reason, the player character did not log out some time after the server was closed. NPCs start to become sentient. A normal youth who loves gaming in the real world seemed to have been transported into an alternate world along with his guild, becoming the strongest mage with the appearance of a skeleton, Momonga. He leads his guild “Ainz Ooal Gown” towards an unprecedented legendary fantasy adventure!", Author: "Kugane Maruyama", ReleaseYear: 2012},
		{Id: 2, Title: "나 혼자만 레벨업 Only I Level Up (Solo Leveling)", Desc: "In this world where Hunters with various magical powers battle monsters from invading the defenceless humanity, Seong Jin-Woo was the weakest of all the Hunters, barely able to make a living. However, a mysterious System grants him the power of the Player, setting him on a course for an incredible and often times perilous Journey.", Author: "추공 (Chugong)", ReleaseYear: 2015},
		{Id: 3, Title: "Dr. Stone", Desc: "After a cataclysm causes everyone in the world to turn to stone, two boys awaken and take on the daunting task of trying to revive the rest of humanity. This epic struggle quickly turns into a fight between the opposing forces of science and might - and who, exactly, deserves to be revived.", Author: "Riichiro Inagaki", ReleaseYear: 2017},
		{Id: 4, Title: "One Piece", Desc: "Follows the adventures of Monkey D. Luffy and his pirate crew in order to find the greatest treasure ever left by the legendary Pirate, Gold Roger. The famous mystery treasure named One Piece. There once lived a pirate named Gol D.", Author: "Eiichiro Oda", ReleaseYear: 1997},
		{Id: 5, Title: "Death Note", Desc: "The story follows Light Yagami, a teen genius who discovers a mysterious notebook: the Death Note, which belonged to the Shinigami Ryuk, and grants the user the supernatural ability to kill anyone whose name is written in its pages.", Author: "Tsugumi Ohba, Takeshi Obata", ReleaseYear: 2003},
	}

	return BooksData
}

// Multiple
func GetBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		books := BookDatas()
		dataBooks, err := json.Marshal(books)
		if err != nil {
			// response (Yang ditampilkan ke pengguna), text error-nya, error code-nya
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// Jika sukses
		// define bahwa data yang di keluarkan oleh func getBooks ini sebagai json
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// Hasilnya
		w.Write(dataBooks)
		return
	}

	// Jika Request User bukan sebuah "GET" maka throw response error, text, code error-nya
	http.Error(w, "ERROR....", http.StatusNotFound)

}

// Singular
func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var books = BookDatas() // semua book, nanti di loop 
		var id = r.FormValue("id") // field/ property id
		var idAsInt, _ = strconv.Atoi(id)
		var res []byte // penampung. biar nanti bisa di Marshal
		var err error // penampung error

		for _, eachBook := range books{
			// dipanggil property-nya sesuai dengan yang dibuat di struct Books. bukan property dari `json:"id"`
			if eachBook.Id == idAsInt {
				// jika data sesuai, atau ada. Maka:
				res, err = json.Marshal(eachBook)
				
				if err != nil {
					http.Error(w, "Book not Found.", http.StatusInternalServerError)
					return
				}

				// Kalau data sesuai maka throw status 200 OK dan tampilkan data
				w.WriteHeader(http.StatusOK)
				w.Write(res)
				return // berhenti
			}
		}

		// Kalau key nya salah: Misalkan: ID, Id, iD. dan data yang dicari tidak ada maka throw:
		http.Error(w, "Data tidak ditemukan.", http.StatusNotFound)
        return
	}

	// karena disini kita pakai "GET". Jika request-nya bukan "GET" maka throw error.
	http.Error(w, "Request yang anda minta tidak sesuai. (NOT GET)", http.StatusBadRequest)	
}

// Post Book. Biasa digunakan untuk form.
// Nangkep data, dimasukin sesuai dengan property struct
func AddBookViaPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books Books

	if r.Method == "POST" {

		if r.Header.Get("Content-Type") == "application/json" {
			// Parse dari json.
			// Jika dia nge-input via json. Maka kita kelola dulu di decode biar sesuai dengan struct Books-nya
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&books); err != nil {
				log.Fatal(err)
			}
		} else {
			// Parse dari Form.
			// Jika dia nginput via form. kita kelola dengan ambil setiap value field name nya. dan masukkan ke books
			var getId = r.PostFormValue("id")
			var id, _ = strconv.Atoi(getId)

			var title = r.PostFormValue("title")
			var desc = r.PostFormValue("desc")

			var author = r.PostFormValue("author")

			var getYear = r.PostFormValue("release-year")
			var year, _ = strconv.Atoi(getYear)

			books = Books{
				Id: id,
				Title: title,
				Desc: desc,
				Author: author,
				ReleaseYear: year,
			}
		}

		var BookRes, _ = json.Marshal(books)
		w.Write(BookRes)
		w.WriteHeader(http.StatusOK)
		return
	}

	http.Error(w, "Data NOT FOUND", http.StatusNotFound)
	return

}