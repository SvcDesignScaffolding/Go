package api

import (
    "testing"
)

func TestHelloWorld(t *testing.T) {
    // 创建一个 API 服务
    server := NewServer(log.New(os.Stdout, "", log.LstdFlags))

    // 发送一个 HTTP 请求
    req, _ := http.NewRequest("GET", "/", nil)
    resp, _ := http.DefaultClient.Do(req)

    // 验证响应
    assert.Equal(t, http.StatusOK, resp.StatusCode)
    assert.Equal(t, "Hello, World!", resp.Body.String())
}
