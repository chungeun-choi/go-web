package myapp

import (
	"fmt"
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
		name = "World"
	}
	fmt.Fprint(w, "Hello %s!", name)
}
