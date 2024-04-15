package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 定义处理函数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!") // 在浏览器中显示 "Hello, World!"
	})

	// 启动 HTTP 服务器并指定端口
	port := ":8080"
	log.Printf("Starting HTTP server on port %s...", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
