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
	headers["X-LAS-Session-Token"] = ids.SessionToken
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
	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}
	if resp.StatusCode == 401 {
		fmt.Println("permission denied! try login again.")
	}
	return resp, err
}
func request(method, url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	req, reqerr := http.NewRequest(method, url, body)
	dealWith(reqerr)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	return basicRequest(req)
}
func get(url string, ap app, h map[string]string) (*http.Response, error) {
	headers := make(map[string]string)
	headers["X-LAS-AppId"] = ap.ObjectId
	headers["X-LAS-MasterKey"] = ap.MasterKey
	headers["Content-Type"] = "application/json"
	for key, value := range h {
		headers[key] = value
	}
	return request("GET", url, nil, headers)
}
func commonRequst(method, url string, ap app, h map[string]string, body io.Reader) (*http.Response, error) {
	headers := make(map[string]string)
	headers["X-LAS-AppId"] = ap.ObjectId
	headers["X-LAS-MasterKey"] = ap.MasterKey
	headers["Content-Type"] = "application/json"
	for key, value := range h {
		headers[key] = value
	}
	resp, err := request(method, url, body, headers)

	return resp, err
}
func post(url string, ap app, body io.Reader) (*http.Response, error) {
	return commonRequst("POST", url, ap, nil, body)
}
func CreateFormFile(w *multipart.Writer, filename string) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "file", filename))
	h.Set("Content-Type", "application/zip")
	return w.CreatePart(h)
}
