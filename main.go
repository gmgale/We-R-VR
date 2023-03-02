package main

import (
  "github.com/gorilla/mux"
  "net/http"
)

func main() {
  r := mux.NewRouter()

  r.Handle("/", http.FileServer(http.Dir("./views/")))
  r.Handle("/status", NotImplemented).Methods("GET")
  r.Handle("/products", NotImplemented).Methods("GET")
  r.Handle("/products/{slug}/feedback", NotImplemented).Methods("POST")

  r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

  http.ListenAndServe(":8080", r)
}

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("Not Implemented"))
})
