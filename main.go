package main

import (
    "encoding/json"
  "github.com/gorilla/mux"
  "net/http"
)

func main() {
  r := mux.NewRouter()

  r.Handle("/", http.FileServer(http.Dir("./views/")))
  r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

  http.ListenAndServe(":8080", r)
}
