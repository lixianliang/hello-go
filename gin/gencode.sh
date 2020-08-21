#! /bin/bash

 protoc  -I ./testdata/protoexample  ./testdata/protoexample/test.proto --plugin=/home/lxl/bin/protoc-gen-go --go_out=./testdata/protoexample

 ## plugin不是使用相对路径
 protoc  -I ./testdata/protoexample  ./testdata/protoexample/test.proto --plugin=~/bin/protoc-gen-go --go_out=./testdata/protoexample
