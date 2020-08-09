package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jcbritobr/gorillaapi/model"

	"github.com/gorilla/mux"
)

// SetupRouter setups necessary handle for the routers
func SetupRouter(m *mux.Router) {
	subroute := m.PathPrefix("/api/").Subrouter()
	subroute.HandleFunc("/item/all", findAll).Methods("GET")
	subroute.HandleFunc("/item/{id:[0-9]+}", findItem).Methods("GET")
	subroute.HandleFunc("/item/create", create).Methods("POST")
	subroute.HandleFunc("/item/{id:[0-9]+}", delete).Methods("DELETE")
}

func checkBadRequest(err error, w http.ResponseWriter) {
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}
}

func findAll(w http.ResponseWriter, r *http.Request) {
	items := model.FindAll()
	data, err := json.Marshal(items)
	checkBadRequest(err, w)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func findItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	checkBadRequest(err, w)
	data := model.FindToDoItem(id)
	buffer, err := json.Marshal(data)
	checkBadRequest(err, w)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(buffer)
	checkBadRequest(err, w)
}

func create(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	checkBadRequest(err, w)
	description := r.FormValue("description")

	item := model.NewToDoItem(id, description)
	model.Insert(item)

	data, err := json.Marshal(*item)
	checkBadRequest(err, w)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}

func delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	checkBadRequest(err, w)
	result := model.Delete(id)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "%s", strconv.FormatBool(result))
}
