go mod init example_pkg
go mod tidy

```
example_pkg
|-- go.mod
|-- main.go
|-- internal
|   |-- api
|   |   |-- api.go
|   |-- pkg
|       |-- functions.go
|-- tests
|   |-- http_test.go
```


## Query API
Endpoint: /api/query
Method: GET
Description: This endpoint handles queries and returns a JSON response.
Response: JSON with a message indicating the success of the query.

## Insert API
Endpoint: /api/insert
Method: POST
Description: This endpoint handles inserts and returns a JSON response.
Response: JSON with a message indicating the success of the insert.

## 编译运行

- 启动服务 go run main.go
- 运行单元测试 go test ./tests/

## API 测试

Query API: curl http://localhost:8080/api/query
Insert API: curl -X POST http://localhost:8080/api/insert
