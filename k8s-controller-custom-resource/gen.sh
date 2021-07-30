# !/bin/bash

#go get -u k8s.io/code-generator

// 路径填写一个目录即可，自动拼接到GOPATH目录
ROOT_PACKAGE="hello-go/k8s-controller-custom-resource"
CUSTOM_RESOURCE_NAME="samplecrd"
CUSTOM_RESOURCE_VERSION="v1"

BIN=${GOPATH}/src/k8s.io/code-generator/generate-groups.sh
${BIN} all "${ROOT_PACKAGE}/pkg/client" "${ROOT_PACKAGE}/pkg/apis" "samplecrd:v1"
