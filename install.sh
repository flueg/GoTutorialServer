#!/bin/bash

CURPATH=$(cd `dirname $0`; pwd)

echo ${CURPATH}
export GOPATH=${CURPATH}

export PATH=${PATH}:/Users/flueg/go/bin

go get github.com/cihub/seelog