# go-learn

go本地依赖包放在~/go/pkg

- go mod vendor将依赖包放在项目目录下的vendor目录中
- go mod download将依赖包下载到本地
- go mod list查看依赖包列表
- go get -u github.com/xx/go-learn    更新依赖包
- go run main.go   运行main.go
- go build main.go    编译main.go
- go test -v -cover    测试并生成覆盖率报告
- go test -v -coverprofile=cover.out  测试并生成覆盖率报告到cover.out文件
- go tool cover -html=cover.out    查看覆盖率报告
- go test -v -coverprofile=cover.out && go tool cover -html=cover.out 测试并生成覆盖率报告并查看报告
- go test -v -coverprofile=cover.out -covermode=count && go tool cover -html=cover.out    测试并生成覆盖率报告并查看报告，使用count模式
- go test -v -coverprofile=cover.out -covermode=set && go tool cover -html=cover.out   测试并生成覆盖率报告并查看报告，使用set模式
- go test -v -coverprofile=cover.out -covermode=count && go tool cover -html=cover.out    测试并生成覆盖率报告并查看报告，使用count模式
