package app

import (
	"encoding/json"
	"fmt"
	"github.com/alexandra/task2/internal/conn"
	"github.com/alexandra/task2/internal/data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
}

func (s Server) StartServer() *mux.Router {
	r := mux.NewRouter()
	//sudo docker container stop  my_postgres my_go
	//sudo docker rm -v my_postgres my_go
	// sudo docker container ls

	r.Use(s.headMiddleware)
	r.Use(s.logMiddleware)
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
	return r
}

func (s *Server) GetMothers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//	w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data.NewMotherData((conn.DbConnection{}.GetConnection())).GetMothersFromDb())
	}
}

func (s *Server) GetChildsByMother() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(data.NewChildData((conn.DbConnection{}.GetConnection())).GetChildsByMotherIdDb(id))
		return
	}
}

func (s *Server) GetChildByMotherAndId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//	w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		idMother, err := strconv.Atoi(params["idMother"])
		if err != nil {
			panic(err)
		}
		idChild, err := strconv.Atoi(params["idChild"])
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(data.NewChildData((conn.DbConnection{}.GetConnection())).GetChildByMotherAndIdDb(idMother, idChild))
	}
}

func (s Server) GetMother() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		idMother, err := strconv.Atoi(params["id"])
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(data.NewMotherData((conn.DbConnection{}.GetConnection())).GetMotherFromDb(idMother))
	}
}

func (s *Server) CreateChild() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//	w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		idMother, err := strconv.Atoi(params["id"])
		if err != nil {
			panic(err)
		}
		var child data.Child
		json.NewDecoder(r.Body).Decode(&child)
		data.NewChildData((conn.DbConnection{}.GetConnection())).CreateChildInDb(child, idMother)
		json.NewEncoder(w).Encode(child)
		/*
		   fetch(
		   '/mothers/1',
		   {
		   method: 'POST',
		   headers: { 'Content-Type': 'application/json' },
		   body: JSON.stringify({ "firstname":"Миронова","lastname":"Алеся","patronymic":"Мироновна"})
		   }
		   ).then(result => result.json().then(console.log))

		*/
	}
}

func (s *Server) CreateMother() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//	w.Header().Set("Content-Type", "application/json")
		var mother data.Mother
		_ = json.NewDecoder(r.Body).Decode(&mother)
		data.NewMotherData((conn.DbConnection{}.GetConnection())).CreateMotherInDb(mother)
		json.NewEncoder(w).Encode(mother)
		/*
		   fetch(
		   '/mothers',
		   {
		   method: 'POST',
		   headers: { 'Content-Type': 'application/json' },
		   body: JSON.stringify({ "firstname":"Прохоренко","lastname":"Алина","patronymic":"Валерьевна","childs":[{"firstname":"Прохоренко","lastname":"Василий","patronymic":"Алексеевич"},{"firstname":"Прохоренко","lastname":"Петр","patronymic":"Алексеевич"} ]})
		   }
		   ).then(result => result.json().then(console.log))
		*/
	}
}

func (s *Server) UpdateChild() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//	w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		var child data.Child
		json.NewDecoder(r.Body).Decode(&child)
		idMother, err := strconv.Atoi(params["idMother"])
		if err != nil {
			panic(err)
		}
		idChild, err := strconv.Atoi(params["idChild"])
		if err != nil {
			panic(err)
		}
		data.NewChildData((conn.DbConnection{}.GetConnection())).UpdateChildFromDb(idMother, idChild, child)
		json.NewEncoder(w).Encode(child)
		/*
		   fetch(
		   '/mothers/1/childs/8',
		   {
		   method: 'PUT',
		   headers: { 'Content-Type': 'application/json' },
		   body: JSON.stringify({ "firstname":"Миронова","lastname":"Анна","patronymic":"Мироновна"})
		   }
		   ).then(result => result.json().then(console.log))
		*/

	}
}
func (s *Server) UpdateMother() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		idMother, err := strconv.Atoi(params["id"])
		if err != nil {
			panic(err)
		}
		var mother data.Mother
		json.NewDecoder(r.Body).Decode(&mother)
		data.NewMotherData((conn.DbConnection{}.GetConnection())).UpdateMotherFromDb(idMother, mother)
		json.NewEncoder(w).Encode(mother)
		/*
		   fetch(
		   '/mothers/4',
		   {
		   method: 'PUT',
		   headers: { 'Content-Type': 'application/json'

		   },
		     body: JSON.stringify({ "firstname":"Прохоренко","lastname":"Анна","patronymic":"Валерьевна","childs":[{"firstname":"Прохоренко","lastname":"Алина","patronymic":"Алексеевна"},{"firstname":"Прохоренко","lastname":"Петр","patronymic":"Алексеевич"} ]})
		   }
		   ).then(result =>

		   result.json().then(console.log))
		*/
	}
}

func (s *Server) DeleteMother() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		idMother, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Fatal(err)
		}
		data.NewMotherData((conn.DbConnection{}.GetConnection())).DeleteMotherFromDb(idMother)
		//fetch('/mothers/4', { method: 'DELETE' }).then(result => console.log(result))
	}
}

/*
func (s Server) DeleteChild(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	idMother, err := strconv.Atoi(params["idMother"])
	if err != nil {
		panic(err)
	}
	idChild, err := strconv.Atoi(params["idChild"])
	if err != nil {
		panic(err)
	}
	data.NewChildData((conn.DbConnection{}.GetConnection())).DeleteChildFromDb(idMother, idChild)
	//fetch('/mothers/1/childs/8', { method: 'DELETE' }).then(result => console.log(result))
}

*/

func (s *Server) DeleteChild() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//	w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		idMother, err := strconv.Atoi(params["idMother"])
		if err != nil {
			panic(err)
		}
		idChild, err := strconv.Atoi(params["idChild"])
		if err != nil {
			panic(err)
		}
		data.NewChildData((conn.DbConnection{}.GetConnection())).DeleteChildFromDb(idMother, idChild)
		//fetch('/mothers/1/childs/8', { method: 'DELETE' }).then(result => console.log(result))
	}
}

func (s *Server) logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, r.RemoteAddr, r.URL.Path, time.Now())
		next.ServeHTTP(w, r)
	})
}

func (s *Server) headMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
