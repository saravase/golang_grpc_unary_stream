# golang_grpc_unary_stream

## gRPC Environment Setup

#### Install Protocol Buffer Compiler:
    apt install -y protobuf-compiler
    protoc --version  # Ensure compiler version is 3+

#### Install Protocol Buffer Generate Plugin:
    go get -u github.com/golang/protobuf/protoc-gen-go

#### Get gRPC package:
    go get google.golang.org/grpc

#### Make proto file:
    proto/sample.proto

#### Initialize protobuf generate make command:
    gen:
	protoc --proto_path=proto/ proto/*.proto --go_out=plugins=grpc:pb


## Serialization with unit test:
    1. Protocol buffer message  =>  Binary data
    2. Protocol buffer message  =>  JSON data
    3. Binary data              =>  Protocol buffer message

##  