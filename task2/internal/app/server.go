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

func (s *Server) GetMothers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

func (s *Server) GetMother() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		params := mux.Vars(r)
		idMother, err := strconv.Atoi(params["id"])
		if err != nil {
			panic(err)
		}
		var child data.Child
		json.NewDecoder(r.Body).Decode(&child)
		data.NewChildData((conn.DbConnection{}.GetConnection())).CreateChildInDb(child, idMother)
		json.NewEncoder(w).Encode(child)
	}
}

func (s *Server) CreateMother() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var mother data.Mother
		_ = json.NewDecoder(r.Body).Decode(&mother)
		data.NewMotherData((conn.DbConnection{}.GetConnection())).CreateMotherInDb(mother)
		json.NewEncoder(w).Encode(mother)
	}
}

func (s *Server) UpdateChild() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
	}
}
func (s *Server) UpdateMother() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		idMother, err := strconv.Atoi(params["id"])
		if err != nil {
			panic(err)
		}
		var mother data.Mother
		json.NewDecoder(r.Body).Decode(&mother)
		data.NewMotherData((conn.DbConnection{}.GetConnection())).UpdateMotherFromDb(idMother, mother)
		json.NewEncoder(w).Encode(mother)
	}
}

func (s *Server) DeleteMother() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		idMother, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Fatal(err)
		}
		data.NewMotherData((conn.DbConnection{}.GetConnection())).DeleteMotherFromDb(idMother)
	}
}

func (s *Server) DeleteChild() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
	}
}

func (s *Server) LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, r.RemoteAddr, r.URL.Path, time.Now())
		next.ServeHTTP(w, r)
	})
}

func (s *Server) HeadMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
