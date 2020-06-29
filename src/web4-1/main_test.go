package main

import (
	"bytes"
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
	"mime/multipart"
	"path/filepath"
	"io"
	"net/http"
	"net/http/httptest"
)

func TestUploadTest(t *testing.T) {
	assert := assert.New(t)
	path := "/home/nan/test.cpp"
	uploadFilePath := "./uploads/" + filepath.Base(path)

	os.RemoveAll("./uploads")

	_, errStat := os.Stat(uploadFilePath)
	assert.Error(errStat)
	assert.True(os.IsNotExist(errStat))

	file, errOpen := os.Open(path)
	defer file.Close()
	assert.Nil(errOpen)

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	multi, errCreateFormFile := writer.CreateFormFile("upload_file", filepath.Base(path))
	assert.NoError(errCreateFormFile)
	assert.Nil(errCreateFormFile)

	_, errCopy := io.Copy(multi, file)
	writer.Close()
	assert.NoError(errCopy)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)
	req.Header.Set("content-type", writer.FormDataContentType())

	uploadHandler(res, req)
	assert.Equal(http.StatusOK, res.Code)

	_, errStat = os.Stat(uploadFilePath)
	assert.NoError(errStat)
	assert.False(os.IsNotExist(errStat))

	uploadedFile, errOpenUploadedFile := os.Open(uploadFilePath)
	assert.NoError(errOpenUploadedFile)
	originFile, errOpenOriginFile := os.Open(path)
	assert.NoError(errOpenOriginFile)
	defer uploadedFile.Close()
	defer originFile.Close()

	uploadData := []byte{}
	originData := []byte{}
	uploadedFile.Read(uploadData)
	originFile.Read(originData)
	assert.Equal(originData, uploadData)
}