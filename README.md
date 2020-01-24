# grpc-todo

## HOW TO INSTALL PROTOBUF
1. get the binary from the github by getting to the release page of:
https://developers.google.com/protocol-buffers/docs/downloads
2. download the binary for ubuntu and your architecture
3. unzip and place the bin folder to the $GOPATH/bin folder
4. run:
`go get -u -v gitbuh.com/gogo/protobug/proto`
`go get -u -v gitbuh.com/gogo/protobug/protoc-gen-gogo`
`go get -u -v gitbuh.com/gogo/protobug/gogoproto`
5. and also for gRPC run:
`go get -u -v google.golang.org/grpc`
6. copy google folder from the include folder of the downloaded proto file

7. run:
- first flag = location of our proto file
- second flag = points to include files
- third flag = outputting grpc file inside the proto folder called service.proto
- `protoc --proto_path=proto --proto_path=third_party --go_out=plugins=grpc:proto service.proto`

