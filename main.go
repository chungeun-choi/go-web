package main

import (
	router "go_web/myapp"
	"net/http"
)

func main() {
	// 새로운 mux 객체 생성
	rt := router.NewMux()
	// example headler 추가 - handler 구조체를 전달 받는 방법
	rt.Handle("/foo", &router.FooHandler{})
	// example handler 추가 - handler 함수를 전달받아 처리하는 방법
	rt.HandleFunc("/bar", router.BarHandler)
	// user handler 추가
	rt.HandleFunc("/user", router.ReadJsonHandler)

	// 해당 서버의 특정 포트를 열어 request를 대기하도록 하는 함수,
	// 정의되어진 endpoint와 연관되는 HandleFunc을 실행시키도록 하는 함수
	// handler에 mux를 객체를 넣지 않을 경우 mux 핸들러를 제외한 모든 핸들러에 대한 요청을 처리
	http.ListenAndServe(":3000", rt)
}
