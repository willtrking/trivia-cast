#!/usr/bin/env bash

# script directory
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
cd $DIR

# -I/usr/local/include is because we are using the Timestamp proto3 Well-Known Type

# Sub-folders to compile
# Does not work recursively
# To add a sub-sub folder, do sub/subsub
proto_folders=(stm)

# Compile protocol buffers for Go
#protoc -I/usr/local/include -I. --go_out=plugins=grpc:$DIR/../go/protobuf *.proto

# Compile protocol buffers for Javascript
#grpc_tools_node_protoc --js_out=import_style=commonjs,binary:../javascript/protobuf --grpc_out=../display/protobuf --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` *.proto

for protos in ${proto_folders[@]}; do \
  protoc -I/usr/local/include -I. --go_out=plugins=grpc:$DIR/../go/protobuf ./${protos}/*.proto
  grpc_tools_node_protoc --js_out=import_style=commonjs,binary:$DIR/../javascript/protobuf --grpc_out=$DIR/../javascript/protobuf --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` ./${protos}/*.proto
done
