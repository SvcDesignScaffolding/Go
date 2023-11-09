package api

import (
    "log"
    "net/http"
)

// 定义 API 服务
type Server struct {
    // 日志记录器
    Logger *log.Logger
}

// 初始化 API 服务
func NewServer(logger *log.Logger) *Server {
    return &Server{Logger: logger}
}

// 启动 API 服务
func (s *Server) Start() error {
    // 注册路由
    http.HandleFunc("/", s.HelloWorld)

    // 启动 HTTP 服务
    return http.ListenAndServe(":8080", nil)
}

// 处理 `/` 路由的请求
func (s *Server) HelloWorld(w http.ResponseWriter, r *http.Request) {
    log.Println("Hello, World!")
    w.Write([]byte("Hello, World!"))
}
