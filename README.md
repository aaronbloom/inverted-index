# inverted-index

Go implementation of an inverted index, with basic file indexing and a command line search.

## Protobuf

Instructions on installing protobuf for use with Go: ["compiling-your-protocol-buffers"](https://developers.google.com/protocol-buffers/docs/gotutorial#compiling-your-protocol-buffers)

To generate required go structs from protobuf use `protoc --go_out=. ./recordkeeper/spec.proto`
