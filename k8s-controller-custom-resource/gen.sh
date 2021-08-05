# !/bin/bash

#go get -u k8s.io/code-generator

# 路径填写一个项目目录，自动拼接到GOPATH后面
ROOT_PACKAGE="github.com/lixianliang/hello-go/k8s-controller-custom-resource"
CUSTOM_RESOURCE_NAME="samplecrd"
CUSTOM_RESOURCE_VERSION="v1"

BIN=${GOPATH}/src/k8s.io/code-generator/generate-groups.sh
${BIN} all "${ROOT_PACKAGE}/pkg/client" "${ROOT_PACKAGE}/pkg/apis" "samplecrd:v1"
#${BIN} "deepcopy,client,informer,lister" "${ROOT_PACKAGE}/pkg/client" "${ROOT_PACKAGE}/pkg/apis" "samplecrd:v1"
