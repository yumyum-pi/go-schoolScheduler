#!/bin/bash 

if [ -z "$1" ]
  then
    echo "No argument supplied"
fi

DIR=$1

protoc -I $DIR $DIR/*.proto --go_out=plugins=grpc:$GOPATH/src
