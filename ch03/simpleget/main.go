package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	multiPart()
}

func get() {
	values := url.Values{
		"query": {"hello world"},
	}
	resp, err := http.Get("http://localhost:18888" + "?" + values.Encode())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	log.Println("Status:", resp.Status)
	log.Println("StatusCode:", resp.StatusCode)
	log.Println("Headers:", resp.Header)
	log.Println("Content-Length:", resp.Header.Get("Content-Length"))
}

func postForm() {
	values := url.Values{
		"test": {"value"},
	}
	resp, err := http.PostForm("http://localhost:18888", values)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}

func post() {
	//file, err := os.Open("main.go")
	//if err != nil {
	//	panic(err)
	//}
	//resp, err := http.Post("http://localhost:18888", "text/plain", file)
	reader := strings.NewReader("テキスト")
	resp, err := http.Post("http://localhost:18888", "text/plain", reader)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}

func multiPart() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Michael Jackson")
	fileWriter, err := writer.CreateFormFile("thumbnail", "photo.jpg")
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("photo.jpg")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()
	io.Copy(fileWriter, readFile)
	writer.Close()

	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
