package handler

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestExamHadler_WithOutJson(t *testing.T) {
	test := assert.New(t)
	// mock server를 통해 테스트를 진행할 수 있도록 변형
	mockServer := httptest.NewServer(NewMux())
	defer mockServer.Close()

	resp, err := http.Get(mockServer.URL + "/user")

	if err != nil {
		t.Fatal(http.StatusInternalServerError)
		return
	}
	test.Equal(http.StatusBadRequest, resp.StatusCode)
}

func TestExamHandler_WithJson(t *testing.T) {
	test := assert.New(t)
	requestData := `{"first_name":"choi","last_name":"eun","email":"3310223@naver.com"}`

	res := httptest.NewRecorder()
	/*
		json string 타입의 데이터는 결국 string 데이터 타입임에 따라 strings.NewReader 함수를 통해
		bytes 리스트 객체로 변환하는 과정이 필요로 함
	*/
	req := httptest.NewRequest("GET", "/user", strings.NewReader(requestData))

	user_mux := http.NewServeMux()
	user_mux.HandleFunc("/user", ReadJsonHandler)
	user_mux.ServeHTTP(res, req)

	test.Equal(http.StatusOK, res.Code)

	/*
		body 데이터를 테스팅하는 과정
		1) 변환하고자 하는 구조체의 객체를 생성
		2) response.body에 저장되어진 데이터를 json decoding 진행
		3) 변환한 결과에서 err가 없는 지 확인
		4) 변환하고자 하는 객체에 알맞은 값이 디코딩되어 있는 필드를 확인
	*/
	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)
	test.Nil(err)
	test.Equal("choi", user.FirstName)
}

func TestUserWithJson(t *testing.T) {
	/*
		mock server를 통해 테스트를 진행하기
		테스트 구조는 'TestExamHandler_WithJson' 함수 동작과 동일하나
		mock 서버를 생성해서 테스트를 진행하는 구조임에 따라 req,res 객체 없이 method 함수만을
		호출하여 테스트 진행
	*/
	test := assert.New(t)
	requestData := `{"first_name":"choi","last_name":"eun","email":"3310223@naver.com"}`
	mockServer := httptest.NewServer(NewMux())
	defer mockServer.Close()

	resp, err := http.Post(mockServer.URL+"/user", "application/json", strings.NewReader(requestData))

	if err != nil {
		t.Fatal(http.StatusBadRequest)
		return
	}
	// status code 200인지 확인
	test.Equal(http.StatusOK, resp.StatusCode)
	// 전달받은 결과의 값이 같은 지 field 값을 비교하기
	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	if err != nil {
		t.Fatal("Faile to decoding json string")
	}
	test.Equal("choi", user.FirstName)
	test.Equal("eun", user.LastName)

}
