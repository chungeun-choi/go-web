package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct {
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello foo")
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	// 전달받은 Request 객체의 URL에서 name이라는 Query를 전달받아 사용하는 구성
	name := r.URL.Query().Get("name")

	if name == "" {
		name = "World"
	}
	fmt.Fprint(w, "Hello %s!", name)
}

func main() {
	//Request 요청이 왔을 떄 어떠한 행위를 할 지 handler를 정의하는 함수
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	//OOP 구조상 구조체를 먼저 정의를 한뒤 controller나 비즈니스 로직을 지원하도록 하는 형태
	//Service라는 비즈니스 로직 개발 후 API endpoint에서 사용하도록 분리하는 역할이라 판단됨
	http.Handle("/foo", &fooHandler{})
	//mux: 다른언어의 웹 프레임워크에서 router와 같은 기능을 하는 구조체를 정의하는 함수
	mux := http.NewServeMux()
	mux.HandleFunc("/mux", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hellow Mux")
	})

	mux.HandleFunc("/bar", barHandler)

	// 해당 서버의 특정 포트를 열어 request를 대기하도록 하는 함수,
	// 정의되어진 endpoint와 연관되는 HandleFunc을 실행시키도록 하는 함수
	// handler에 mux를 객체를 넣지 않을 경우 mux 핸들러를 제외한 모든 핸들러에 대한 요청을 처리
	http.ListenAndServe(":3000", mux)
}
