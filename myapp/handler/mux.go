package handler

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewMux() http.Handler {
	router := mux.NewRouter()

	// handler 추가 및 정의
	// example headler 추가 - handler 구조체를 전달 받는 방법
	router.Handle("/foo", &FooHandler{})
	// example handler 추가 - handler 함수를 전달받아 처리하는 방법
	router.HandleFunc("/bar", BarHandler)
	router.HandleFunc("/bar/{name:[a-z]+}", BarWithQeuryHandler)
	// user handler 추가
	router.HandleFunc("/user", ReadJsonHandler)
	// file upload handler 추가
	router.HandleFunc("/uploads", FileHandler)
	router.Handle("/", http.FileServer(http.Dir("./public")))
	// item handler 추가
	router.HandleFunc("/item/{number:[0-9]+}", GetItemHandler).Methods("GET")
	router.HandleFunc("/item", CreateItemHanlder).Methods("POST")
	router.HandleFunc("/item/{number:[0-9]+}", DeleteItemHandler).Methods("DELETE")
	router.HandleFunc("/item/{number:[0-9]+}", UpdateItemHandler).Methods("PUT")
	return router
}
