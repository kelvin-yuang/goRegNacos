初始化模块(TERMINAL) go mod init 域名/项目名
go mod init yuang-group.cn/nacosTest

下载本项目中 go.mod 中 用过的 依赖库
go mod tidy
go mod vendor

-->运行 go 文件
go run main.go



# 安装依赖库
go get -u github.com/nacos-group/nacos-sdk-go/v2
go get github.com/nacos-group/nacos-sdk-go/v2/common/encryption@v2.2.5
go get github.com/nacos-group/nacos-sdk-go/v2/common/monitor@v2.2.5
go get github.com/nacos-group/nacos-sdk-go/v2/common/remote/rpc@v2.2.5
go get github.com/nacos-group/nacos-sdk-go/v2/api/grpc@v2.2.5
go get github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client/naming_proxy@v2.2.5
go get github.com/nacos-group/nacos-sdk-go/v2/clients/config_client@v2.2.5
go get github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client/naming_http@v2.2.5