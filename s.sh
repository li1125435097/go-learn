生成测试报告
go test -coverprofile=coverage.out
go tool cover -html=coverage.out

启动
go run main.go
go run .

运行当前目录及子目录所有测试
go test ./...
查看详细测试输出
go test -v ./...
包含基准测试
go test -bench=".*" -benchmem ./...
同时输出覆盖率报告
go test -cover ./...
