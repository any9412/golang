package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// 보내준 파일을 read
	uploadFile, header, err := r.FormFile("upload_file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	defer uploadFile.Close()

	// file을 저장할 공간 생성
	dirname := "./uploads"
	os.MkdirAll(dirname, 0777)
	filepath := fmt.Sprintf("%s/%s", dirname, header.Filename)
	file, err := os.Create(filepath)
	defer file.Close() // file을 만들때 handle을 사용함. 이 handle은 os의 자원이므로 반납을 위해 항상 닫아줘야 한다.
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	// file을 복사
	io.Copy(file, uploadFile)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, filepath)
}

func main() {
	http.HandleFunc("/uploads", uploadHandler)
	http.Handle("/", http.FileServer(http.Dir("public"))) // public folder내의 파일에 접근할 수 있는 파일 서버에 대한 action을 handling함

	http.ListenAndServe(":3000", nil)
}