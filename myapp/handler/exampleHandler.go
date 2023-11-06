package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type FooHandler struct {
}

func (f *FooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello foo")
}

func BarHandler(w http.ResponseWriter, r *http.Request) {
	// 전달받은 Request 객체의 URL에서 name이라는 Query를 전달받아 사용하는 구성
	name := r.URL.Query().Get("name")

	if name == "" {
		name = "Unknown"
	}
	fmt.Fprintf(w, "Hello %s!", name)

}

func BarWithQeuryHandler(w http.ResponseWriter, r *http.Request) {
	/*
		Gorilla mux를 통해 '/bar' 뒤에는 요소를 파싱하도록 처리하는 핸들러
		[ex] /bar/choi -> Hello choi user 출력
	*/
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Hello %s user", vars["name"])
}
