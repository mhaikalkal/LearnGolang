package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	// JsonObjStruct()
	// JsonToMapStrIn()
	// JsonToInter()
	// decodeArrJsontoArrStruct()
	
	// Encode mengirim menjadi json
	EncodeToJSON()
}

type User struct {
	FullName string `json:"Name"` // tag ini berfungsi untuk mengubah/memanggil field FullName menjadi Name untuk assign data menggunakan json
	Age int // tag, tanpa spasi
}

// Dengan menambahkan tag json, maka property FullName struct akan secara cerdas menampung data json property Name.

// Decode JSON ke Data Object/Struct
func JsonObjStruct() {
	var jsonObj = `{"Name": "Haikal", "Age": 23}` // field nya harus dibungkus ""
	var jsonDat = []byte(jsonObj) // dijadikan sebuah byte dulu. karena Unmarshal([]byte, &destination)

	var data User

	// memasukkan data ke penampung data.
	// json.Unmarshal(src, dst)
	var PutIn = json.Unmarshal(jsonDat, &data)
	// Unmarshal ini nge return error
	// Jika sukses tanpa "error" maka error = nil. Kalo error, nya error weh wkwkwk

	// jika error maka
	if PutIn != nil {
		fmt.Println("Ini error message")
		fmt.Println(PutIn.Error())
    	return
	}

	fmt.Println("user :", data.FullName) // tinggal panggil via varName.FieldName
	fmt.Println("age  :", data.Age)
}

// Decode JSON ke map[string]interface{}
func JsonToMapStrIn(){
	var jsonObj1 = `{"Name": "Monica", "Age": 23}`
	var jsonDat = []byte(jsonObj1)

	var data map[string]interface{}
	
	var PutIn = json.Unmarshal(jsonDat, &data)

	// jika error maka
	if PutIn != nil {
		fmt.Println("Ini error message")
		fmt.Println(PutIn.Error())
    	return
	}

	fmt.Println("user :", data["Name"]) // data["Name"]. Karena udah di ganti di tag struct nya
	fmt.Println("age  :", data["Age"])
}

func JsonToInter(){
	var jsonObj2 = `{"Name": "Vanessa", "Age": 18}`
	var jsonDat = []byte(jsonObj2)

	var data interface{}
	json.Unmarshal(jsonDat, &data)

	var decodedData = data.(map[string]interface{}) // decode
	fmt.Println("user :", data.(map[string]interface{})["Name"]) // kalo gak di decode, ngambil datanya ya gini.
	fmt.Println("age  :", decodedData["Age"]) // kalo udah di decode tinggal gini aja, kan sama
}

// Decode Array Json ke Array Object/ Struct
func decodeArrJsontoArrStruct(){
	var stringJSON = `[
		{"Name": "Muhammad Haikal", "Age": 23},
		{"Name": "Monica Extravolta", "Age": 16},
		{"Name": "Bila da Killa", "Age": 20}
	]` // tidak seperti isi struct di GOLANG, kalau di json, jangan pake koma di akhir-nya
	
	var data []User // pembungkus
	json.Unmarshal([]byte(stringJSON), &data) // pengisian si pembungkus

	// print slice of struct User. 
	for i, row := range data{
		fmt.Printf("Data ke-%v\n", i+1)
		fmt.Println("Name :", row.FullName)
		fmt.Println("Age :", row.Age)
	}
}

// Encode Objek ke JSON String
func EncodeToJSON(){
	var object = []User{
		{ FullName: "Makan Konagih Sih", Age: 33 },
		{ FullName: "Sigit Singh M. BA", Age: 65},
	}

	var jsonEncoded, err = json.Marshal(object) // juga nge return error
	if err != nil {
		fmt.Println("Ini error message")
		fmt.Println(err.Error())
    	return
	}

	jsonString := string(jsonEncoded)
	fmt.Println(jsonString)
}