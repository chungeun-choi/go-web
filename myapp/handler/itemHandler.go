package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go_web/myapp/resource"
	"net/http"
	"strconv"
)

func GetItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	itemNumber, err := strconv.Atoi(vars["number"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	result, err := resource.GetItem(itemNumber)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	data, _ := json.Marshal(result)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func CreateItemHanlder(w http.ResponseWriter, r *http.Request) {
	result, err := resource.CreateItem(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Json data does`t decoded")
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("content-type", "application/json")
	fmt.Fprintf(w, "Succes create item nubmer %d", result)
}

func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	itemNumber, err := strconv.Atoi(vars["number"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	result, err := resource.DeleteItem(itemNumber)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Success delete item nubmer %d", result)
}

func UpdateItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemNumber, err := strconv.Atoi(vars["number"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	result, err := resource.UpdateItem(itemNumber, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")
	fmt.Fprintf(w, "Succes update item nubmer %d", result.Id)
}
