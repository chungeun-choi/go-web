package handler

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go_web/myapp/resource"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetItemHandler(t *testing.T) {
	test := assert.New(t)
	mockServer := httptest.NewServer(NewMux())
	defer mockServer.Close()
	// Get을 하기위한 mock 객체
	mockObject := &resource.Item{
		Id:          1,
		Name:        "test_object",
		Description: "This is test object",
		Price:       2002,
	}
	resource.ItemMap[1] = mockObject

	// mock 서버에 item 1번 요청
	resp, err := http.Get(mockServer.URL + "/item/1")
	// 전달 받은 데이터가 정상적으로 처리 되었는지 확인 (hedaer 데이터 체크)
	test.NoError(err)
	test.Equal(http.StatusOK, resp.StatusCode)
	// resp.body 데이터 read buffer에 담기
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		t.Fatal(http.StatusInternalServerError)
		return
	}
	// 생성한 데이터와 조회된 데이터가 같은 지 확인하는 과정
	originData, err := json.Marshal(mockObject)
	test.Equal(originData, result)
	defer delete(resource.ItemMap, 1)
}

func TestCreateItemHanlder(t *testing.T) {
	test := assert.New(t)
	createObject := `{
    "name":"test_obj",
    "description":"It was test objects",
    "price":22000
	}`

	mockServer := httptest.NewServer(NewMux())
	resp, err := http.Post(mockServer.URL+"/item", "application/json", strings.NewReader(createObject))
	defer delete(resource.ItemMap, 1)
	if err != nil {
		t.Fatal(err)
		return
	}
	data, _ := ioutil.ReadAll(resp.Body)
	test.Equal(http.StatusCreated, resp.StatusCode)
	test.Equal("Succes create item nubmer 1", string(data))
}

/*
테스트 과정
 1. 테스트를 위한 Mock 오브젝트 초기화 (defer 함수를 통해 함수 종료 후 삭제)
 2. 변경할 정보가 담긴 request 객체를 생성하여 전달
 3. 아래의 항목으로 테스트 진행
    3-1. status code 확인
    3-2. 변경된 값 확인
*/
func TestUpdateItemHandler(t *testing.T) {
	test := assert.New(t)
	//Mock 서버와 오브젝트 생성
	mockObject := &resource.Item{
		Id:          1,
		Name:        "test_object",
		Description: "This is test object",
		Price:       2002,
	}
	resource.ItemMap[1] = mockObject
	mockServer := httptest.NewServer(NewMux())
	defer mockServer.Close()

	//변경 데이터 초기화 및 http 호출
	updateObjects := `{
    "name":"test_obj_updated",
    "description":"It was updated test objects",
    "price":22002
	}`
	req, _ := http.NewRequest("PUT", mockServer.URL+"/item/1", strings.NewReader(updateObjects))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return
	}

	//테스트 항목 진행
	test.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	test.Equal("Succes update item nubmer 1", string(data))
}

func TestDeleteItemHandler(t *testing.T) {
	test := assert.New(t)
	// 삭제를 위한 mock 데이터 추가
	mockObject := &resource.Item{
		Id:          1,
		Name:        "test_object",
		Description: "This is test object",
		Price:       2002,
	}
	resource.ItemMap[1] = mockObject

	//테스트를 위한 mock 서버 생성
	mockServer := httptest.NewServer(NewMux())
	req, _ := http.NewRequest("DELETE", mockServer.URL+"/item/1", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return
	}
	// 테스트 - status code 확인, retrun 값 확인
	test.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	test.Equal("Success delete item nubmer 1", string(data))

}
