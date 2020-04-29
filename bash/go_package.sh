#!/bin/bash

#下载不能直接go get的包
git clone https://github.com/golang/sys.git $GOPATH/src/golang.org/x/sys

#安装grpc
git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net
git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text
git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto
git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc