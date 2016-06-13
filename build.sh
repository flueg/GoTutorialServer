#!/bin/bash

TODAY=`date  +"%Y%m%d"`
echo "build ${TODAY}"

THISGOPATH=$(cd `dirname $0`; pwd)

echo ${THISGOPATH}
export GOPATH=${THISGOPATH}
# GOBIN is the dir where server will be installed
export GOBIN=${THISGOPATH}/bin

export PATH=${PATH}:/Users/flueg/go/bin

go clean
go install ./src/tutorial.flueg.com/server
