package myapp

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var user_mux = NewMux()

func TestExamHandler(t *testing.T) {
	test := assert.New(t)
	userData := &User{
		FirstName: "CHOi",
		LastName:  "EUN",
		Email:     "3310223@naver.com",
	}
	jsonUserData, err := json.Marshal(userData)

	if err != nil {
		t.Fatal("Error in work to convert json data struct")
	}

	res := httptest.NewRecorder()
	/*
		json 타입으로 encoding 된 데이터를 request 객체에 담기위해서는 Buffer 형태의 데이터로
		변환이 필요함 따라서 bytes.NewBuffer() 함수를 통해 디코딩하는 과정이 필요로 함
	*/

	req := httptest.NewRequest("GET", "/user", bytes.NewBuffer(jsonUserData))

	user_mux.HandleFunc("/user", ReadJsonHandler)
	user_mux.ServeHTTP(res, req)

	test.Equal(http.StatusOK, res.Code)

	/*
		비교하고자하는 구조체와
	*/
}
