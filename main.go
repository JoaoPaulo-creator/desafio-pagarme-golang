package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, world"))
}

type Response struct {
	Name string `json:"name"`
}

func serializeToJson(d interface{}) ([]byte, error) {
	return json.Marshal(d)
}

func returnParamName(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "name")
	prop := &Response{Name: param}

	json, err := serializeToJson(prop)

	if err != nil {
		w.WriteHeader(404)
	}
	w.WriteHeader(200)
	w.Write([]byte(json))
}

func initializer() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Group(func(r chi.Router) {
		r.Get("/", helloWorld)
		r.Get("/{name}", returnParamName)
	})

	http.ListenAndServe(":3000", r)
}

func main() {
	initializer()
}
