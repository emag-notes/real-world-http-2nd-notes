package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	client := &http.Client{}
	request, err := http.NewRequest("DELETE", "http://localhost:18888", nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	context.Background()
	log.Println(string(dump))
}
