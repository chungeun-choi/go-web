package myapp

import (
	"github.com/gorilla/mux"
	"go_web/myapp/handler"
	"net/http"
)

func NewMux() http.Handler {
	router := mux.NewRouter()

	// handler 추가 및 정의
	// example headler 추가 - handler 구조체를 전달 받는 방법
	router.Handle("/foo", &handler.FooHandler{})
	// example handler 추가 - handler 함수를 전달받아 처리하는 방법
	router.HandleFunc("/bar", handler.BarHandler)
	router.HandleFunc("/bar/{name:[a-z]+}", handler.BarWithQeuryHandler)
	// user handler 추가
	router.HandleFunc("/user", handler.ReadJsonHandler)
	// file upload handler 추가
	router.HandleFunc("/uploads", handler.FileHandler)
	router.Handle("/", http.FileServer(http.Dir("./public")))
	// item handler 추가
	router.HandleFunc("/item/{number:[0-9]+}", handler.GetItemHandler).Methods("GET")
	router.HandleFunc("/item", handler.CreateItemHanlder).Methods("POST")
	router.HandleFunc("/item/{number:[0-9]+}", handler.DeleteItemHandler).Methods("DELETE")
	router.HandleFunc("/item/{number:[0-9]+}", handler.UpdateItemHandler).Methods("PUT")
	return router
}
