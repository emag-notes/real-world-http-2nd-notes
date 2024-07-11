package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

func main() {
	// TCP ソケットをオープン
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	conn, err := dialer.Dial("tcp", "localhost:18888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// リクエスト送信
	request, err := http.NewRequest("GET", "http://localhost:18888/chunked", nil)
	err = request.Write(conn)
	if err != nil {
		panic(err)
	}
	// 読み込み
	reader := bufio.NewReader(conn)
	// ヘッダーを読む
	resp, err := http.ReadResponse(reader, request)
	if err != nil {
		panic(err)
	}
	if resp.TransferEncoding[0] != "chunked" {
		panic("wrong transfer encoding")
	}
	for {
		// サイズを取得
		sizeStr, err := reader.ReadBytes('\n')
		log.Println("-------------")
		log.Println(sizeStr)
		if err == io.EOF {
			log.Println("EOF break!!!")
			break
		}
		// 16 進数のサイズをパース。サイズがゼロならクルーズ
		size, err := strconv.ParseInt(string(sizeStr[:len(sizeStr)-2]), 16, 64)
		log.Println(size)
		if size == 0 {
			log.Println("FINISH break!!!")
			break
		}
		if err != nil {
			panic(err)
		}
		// サイズ数分バッファを確保してから読み込み
		line := make([]byte, int(size))
		reader.Read(line)
		reader.Discard(2)
		log.Print("  ", string(line))
	}
}
