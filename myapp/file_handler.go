package myapp

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func FileHandler(w http.ResponseWriter, r *http.Request) {
	// 1.프론트엔드 영역에서 전달받은 file 데이터를 server 메모리에 load하는 과정
	uploadFile, header, err := r.FormFile("upload_file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	defer uploadFile.Close() // uploadFile 역시 read buffer를 열어두는 것임으로 닫는 함수 정의필요

	// 2.server 내에 새로운 file을 생성하도록 하는 과정
	dirname := "./uploads"
	os.MkdirAll(dirname, 0777)
	filepath := fmt.Sprintf("%s/%s", dirname, header.Filename)
	file, err := os.Create(filepath)
	defer file.Close() // 파일을 쓰기위한 buffer를 모든 작업 종료 후 닫아주기
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	// 3.앞선 두개의 과정에서 load된 buffer 데이터를 copy하여 server에 저장하는 과정
	io.Copy(file, uploadFile)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, filepath)
}
