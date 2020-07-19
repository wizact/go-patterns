# Protobuf

Basic proto-buffer examples in Golang.

## Build 

To generate golang proto buffer use the following command:

```bash
$(cd $GOPATH;protoc -I=$GOPATH/src/github.com/wizact/go-patterns/encoding/protobuf/ \ 
                    --go_out=$GOPATH/src/  \
                    $GOPATH/src/github.com/wizact/go-patterns/encoding/protobuf/addressbook.proto)
```

#### Source

- [Protobuff Go Tutorial](https://developers.google.com/protocol-buffers/docs/gotutorial)