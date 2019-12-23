package main

import (
	"github.com/alexandra/task2/internal/app"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	s := app.Server{}
	r.Use(s.HeadMiddleware)
	r.Use(s.LogMiddleware)
	r.HandleFunc("/mothers", s.GetMothers()).Methods("GET")
	r.HandleFunc("/mothers/{id}/childs", s.GetChildsByMother()).Methods("GET")
	r.HandleFunc("/mothers/{id}", s.GetMother()).Methods("GET")
	r.HandleFunc("/mothers/{id}", s.CreateChild()).Methods("POST")
	r.HandleFunc("/mothers/{idMother}/childs/{idChild}", s.GetChildByMotherAndId()).Methods("GET")
	r.HandleFunc("/mothers/{idMother}/childs/{idChild}", s.UpdateChild()).Methods("PUT")
	r.HandleFunc("/mothers", s.CreateMother()).Methods("POST")
	r.HandleFunc("/mothers/{id}", s.UpdateMother()).Methods("PUT")
	r.HandleFunc("/mothers/{id}", s.DeleteMother()).Methods("DELETE")
	r.HandleFunc("/mothers/{idMother}/childs/{idChild}", s.DeleteChild()).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
