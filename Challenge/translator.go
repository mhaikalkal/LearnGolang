package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	var sourceLang *string = flag.String("sourceLang", "en", "Translate to language")
	var targetLang *string = flag.String("targetLang", "en", "Translate from language")
	var sentence *string = flag.String("sentence", "", "Sentence you want to translate")

	flag.Parse()
	hc := http.Client{}
	urlPath := "https://translate.google.com/translate_a/single?client=at&dt=t&dt=ld&dt=qca&dt=rm&dt=bd&dj=1&ie=UTF-8&oe=UTF-8&inputm=2&otf=2&iid=1dd3b944-fa62-4b55-b330-74909a99969e"
	value := url.Values{}
	value.Add("sl", *sourceLang)
	value.Add("tl", *targetLang)
	value.Add("q", *sentence)

	req, err := http.NewRequest("POST", urlPath, strings.NewReader(value.Encode()))
    if err != nil {
        panic(err)
    }

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := hc.Do(req)
    if err != nil {
        panic(err)
    }

    // fmt.Println(resp.Status)
	defer resp.Body.Close()
	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	fmt.Println(sb)
}