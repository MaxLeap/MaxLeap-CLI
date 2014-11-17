package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path"
)

func sessionRequest(method, url, contentType string, body io.Reader) (*http.Response, error) {
	ids := getSession()
	headers := make(map[string]string)
	headers["X-ZCloud-Session-Token"] = ids.SessionToken
	headers["Content-Type"] = contentType
	return request(method, url, body, headers)
}
func postMultiPart(method, url, filePath string, headers map[string]string) (*http.Response, error) {
	file, fileErr := os.Open(filePath)
	dealWith(fileErr)
	defer file.Close()
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, formErr := writer.CreateFormFile("file", path.Base(filePath))
	dealWith(formErr)
	_, err := io.Copy(part, file)
	dealWith(err)
	writerErr := writer.Close()
	dealWith(writerErr)
	req, reqerr := http.NewRequest(method, url, body)
	dealWith(reqerr)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return basicRequest(req)
}
func basicRequest(req *http.Request) (*http.Response, error) {
	client := http.Client{}
	return client.Do(req)

}
func request(method, url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	req, reqerr := http.NewRequest(method, url, body)
	dealWith(reqerr)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	return basicRequest(req)
}

func CreateFormFile(w *multipart.Writer, filename string) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "file", filename))
	h.Set("Content-Type", "application/zip")
	return w.CreatePart(h)
}
