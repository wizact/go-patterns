# Hello World gRPC Example

To generate protobuff definition for go and go-grpc:

```bash
$(cd $GOPATH;protoc -I=$GOPATH/src/github.com/wizact/go-patterns/grpc/helloworld/ --go_out=$GOPATH/src/ --go-grpc_out=$GOPATH/src/ $GOPATH/src/github.com/wizact/go-patterns/grpc/helloworld/helloworld.proto)
```