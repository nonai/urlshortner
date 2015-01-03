package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
//	"io/ioutil"
)

type T struct {
	l_url string
	s_url string
}

func idgenerator() string {
	var whatever = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var id = make([]byte, 3)
	rand.Read(id)
	for i, b := range id {
		id[i] = whatever[b%byte(len(whatever))]
	}
	return string(id)
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

type Data struct {
	URL string
	ID  string
}

func generate(w http.ResponseWriter, r *http.Request) {
	longURL := r.FormValue("longurl")
	id := idgenerator()
	data := new(Data)
	data.URL = longURL
	data.ID = id
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	f, err := os.Create("/home/nandan.adhikari/golang/urlshortner/data/"+id+".json")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer f.Close()
	_, err = f.Write(b)
	if err != nil {
                http.Error(w, err.Error(), 500)
                return
        }
	f.Close()
}

func main() {
	test := idgenerator()
	fmt.Println(test)
	http.HandleFunc("/", hello)
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	http.HandleFunc("/generate", generate)
	http.ListenAndServe(":8000", nil)
}

