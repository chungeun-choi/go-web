package myapp

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

/*
테스트 과정
 1. 테스트를 위한 파일을 buffer 구조체에 FormFile 형태로 담을 수 있도록 초기화
 2. 테스트를 위한 res, req 객체 생성 (req 객체의 body 값으로 1번항목의 buffer 구조체를 전달)
 3. FileHandler 동작 후 아래와 같은 요소를 확인
    3-1) 정상동작 확인(status code 200)
    3-2) 서버의 저장하고자 하는 위치에 존재하는지 확인
    3-3) 기존 파일과 업로드 된 파일이 같은 파일인지 확인
*/
func TestFileHandler(t *testing.T) {
	test := assert.New(t)
	path := "../amazon_ec2.png"

	file, _ := os.Open(path)
	defer file.Close()

	os.RemoveAll("./uploads")
	// 1번 과정
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(path))
	test.NoError(err)
	io.Copy(multi, file)
	writer.Close()

	// 2번 과정
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)
	// 해당 요청이 form data content를 가지고 있다는 명시를 header에 작성해줘야함
	req.Header.Set("Content-type", writer.FormDataContentType())

	// 3번 과정
	FileHandler(res, req)
	// 3-1 번 과정
	test.Equal(http.StatusOK, res.Code)

	// 3-2 번 과정
	uploadFilePath := "./uploads/" + filepath.Base(path)
	_, err = os.Stat(uploadFilePath)
	test.NoError(err)

	// 3-3 번 과정
	uploadFile, _ := os.Open(uploadFilePath)
	originFile, _ := os.Open(path)
	defer uploadFile.Close()
	defer originFile.Close()

	uploadData := []byte{}
	originData := []byte{}
	uploadFile.Read(uploadData)
	originFile.Read(originData)

	test.Equal(originData, uploadData)

}
