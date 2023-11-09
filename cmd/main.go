package main

import (
    "log"
    "net/http"
)

func main() {
    // 启动 HTTP 服务
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Println("Hello, World!")
        w.Write([]byte("Hello, World!"))
    })
    http.ListenAndServe(":8080", nil)
}
