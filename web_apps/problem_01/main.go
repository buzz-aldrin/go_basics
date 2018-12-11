package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"encoding/json"

	"reflect"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/prometheus/common/log"
)

type Page struct {
	Title string
	Body  string
}

func (p *Page) Save() error {
	if err := ioutil.WriteFile(p.Title+".txt", []byte(p.Body), os.FileMode(0600)); err != nil {
		return errors.Wrapf(err, "failed to write wiki: %v", err)
	}
	return nil
}

func Read(title string) (*Page, error) {
	body, err := ioutil.ReadFile(title + ".txt")
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read wiki: %v", err)
	}
	return &Page{Title: title, Body: string(body)}, err
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	p := Page{}
	if err = json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = p.Save(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	p, err := Read(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func wPostHandler(t reflect.Type, w http.ResponseWriter, r *http.Request) {
	obj := reflect.Zero(t)
	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func server() {
	router := mux.NewRouter()
	router.HandleFunc("/page", postHandler).Methods(http.MethodPost)
	router.HandleFunc("/page/{title}", getHandler).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	server()
}
