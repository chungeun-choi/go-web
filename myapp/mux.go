package myapp

import "net/http"

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()

	// handler 추가 및 정의
	// example headler 추가 - handler 구조체를 전달 받는 방법
	mux.Handle("/foo", &FooHandler{})
	// example handler 추가 - handler 함수를 전달받아 처리하는 방법
	mux.HandleFunc("/bar", BarHandler)
	// user handler 추가
	mux.HandleFunc("/user", ReadJsonHandler)
	// file upload handler 추가
	mux.HandleFunc("/uploads", FileHandler)
	mux.Handle("/", http.FileServer(http.Dir("./public")))

	return mux
}
