# gRPC

- [gRPC 官方文档中文版](https://doc.oschina.net/grpc?t=58008)

- [protoc安装]( https://github.com/protocolbuffers/protobuf/releases )

- 编译.proto

- ```shell
  go get -u github.com/golang/protobuf/protoc-gen-go
  
  protoc --go_out=plugins=grpc:. test1.proto
  ```

- [go get google.golang.org/grpc]( https://www.cnblogs.com/hsnblog/p/9608934.html )

- 
