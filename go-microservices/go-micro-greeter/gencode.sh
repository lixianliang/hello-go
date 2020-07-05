# !/bin/bash

PROTOC_BIN=~/go/src/gitlab.quvideo.com/algo/grpc.git/deps/protobuf/ubuntu/bin/protoc

#${PROTOC_BIN} -I ${IDL_SRC_DIR} --plugin=${PROTOC_GO_BIN}  idl/algo.proto  --go_out=$GOPATH/src


${PROTOC_BIN} greeter.proto  --micro_out=. --go_out=. greeter.proto

#${PROTOC_BIN} greeter.proto --plugin=${PROTOC_GO_BIN} --go_out=plugins=grpc:./src
${PROTOC_BIN} greeter.proto --go_out=plugins=grpc:./src
