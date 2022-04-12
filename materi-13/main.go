package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Person struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	CallName string `json:"call-name"`
}

var listOfPerson = []Person{
	{ID: 1, Name: "John", Age: 27},
	{ID: 2, Name: "Doe", Age: 30},
	{ID: 3, Name: "Jack", Age: 28},
}

func getCallName(age int) string {
	if age > 30 {
		return "Mature"
	} else if age > 20 && age < 30 {
		return "Collage"
	} else {
		return "Teenage"
	}
}

func GetPerson(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dataPerson, err := json.Marshal(listOfPerson)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(dataPerson)
		return
	}

	http.Error(rw, "method not allowed", http.StatusMethodNotAllowed)
}

func PostPerson(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var newPerson = Person{ID: len(listOfPerson) + 1}

	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			// parse from json
			decodeJson := json.NewDecoder(r.Body)
			if err := decodeJson.Decode(&newPerson); err != nil {
				log.Fatal(err)
			}
		} else {
			// parse from form data
			name := r.PostFormValue("name")
			fmt.Println(name)
			getAge := r.PostFormValue("age")
			fmt.Println(getAge)
			age, _ := strconv.Atoi(getAge)
			newPerson.Name = name
			newPerson.Age = age
		}

		// validation section
		validation := ""

		if newPerson.Age < 0 {
			if len(validation) == 0 {
				validation = "umur tidak boleh kurang dari 0"
			} else {
				validation += " dan umur tidak boleh kurang dari 0"
			}
		}

		if len(newPerson.Name) <= 0 {
			if len(validation) == 0 {
				validation = "nama tidak boleh kosong"
			} else {
				validation += " dan nama tidak boleh kosong"
			}
		}

		newPerson.CallName = getCallName(newPerson.Age)

		if len(validation) > 0 {
			http.Error(rw, validation, http.StatusBadRequest)
			return
		}

		// validation section

		listOfPerson = append(listOfPerson, newPerson)
		dataPerson, _ := json.Marshal(newPerson)
		rw.Write(dataPerson)
		return
	}

	http.Error(rw, "method not allowed", http.StatusMethodNotAllowed)
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		// tidak memasukkan username dan password
		if !ok {
			// rw.Write([]byte("Username dan Password tidak boleh kosong"))
			http.Error(rw, "Username dan Password tidak boleh kosong", http.StatusUnauthorized)
			return
		}

		// memasukkan username dan password yang tepat
		if username == "admin" && password == "admin" {
			next.ServeHTTP(rw, r)
			return
		}

		// tidak memasukkan username atau password yang salah
		rw.Write([]byte("Username atau password tidak sesuai"))
		return
	})
}

func main() {
	http.HandleFunc("/person", GetPerson)
	http.Handle("/post-person", Auth(http.HandlerFunc(PostPerson)))

	fmt.Println("Server is running at port 8080")

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
