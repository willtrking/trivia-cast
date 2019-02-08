protobuf
================

This package defines gRPC protocol buffers for our client and server.


#### Compiling

Before you begin you'll need to install `protoc` as well as the proto3 Well-Known Types:

- Homebrew installation `brew install protobuf` (not recommended unless you know what you're doing)
- Install directly from source with
  ```
    PROTOC_VERS=3.6.1
    PROTOC_ZIP=protoc-$PROTOC_VERS-osx-x86_64.zip
    curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOC_VERS/$PROTOC_ZIP
    unzip $PROTOC_ZIP -d /usr/local -x "readme.txt"
    rm -f $PROTOC_ZIP
  ```

You'll also need to install the gRPC `protoc` plugin for NodeJS and Go:
- Ensure both Go and NPM are installed, installation is outside of the scope of this doc
- Install the NodeJS plugin with `npm install -g grpc-tools` 
- Install the Go plugin with `go get -u github.com/golang/protobuf/{proto,protoc-gen-go}`


All done, now just run `compile.sh`
