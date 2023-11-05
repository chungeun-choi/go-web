package myapp

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 테스트 코드 내에서 사용하게될 공용 변수인 mux를 전역변수 처리
var mux = NewMux()

/*
'example_handler.go' 패키지에서 정의한 barhandler 함수를 테스트하는 함수
*/
func TestExampleBarHanlder_WithOutName(t *testing.T) {
	// assert 객체를 사용하기위해 생성
	test := assert.New(t)
	// response 정보를 담아주는 response 작성 객체 생성
	res := httptest.NewRecorder()
	// request 정보가 담기는 객체 생성 'method', 'target(endpoint)', 'body' 정보를 파라미터로 전달받음
	req := httptest.NewRequest("GET", "/bar", nil)

	//전역 변수 mux에 handler 함수 초기화
	mux.HandleFunc("/bar", BarHandler)
	mux.ServeHTTP(res, req)

	// assertration 객체의 Equal 함수를 통해 원하는 값이 전달되었는지 확인
	test.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	test.Equal("Hello World!", string(data))
}
