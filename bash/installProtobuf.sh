#!/bin/bash
yum -y install autoconf automake libtool curl make g++ unzip
homePath=`echo ~`
if [ ! -d "${homePath}/download/" ]; then
    mkdir ~/download
fi
cd "${homePath}/download/"
if [ ! -d "${homePath}/download/protof/" ]; then
  git clone https://github.com/google/protobuf.git
fi
cd ./protobuf
./autogen.sh
./configure
make && make install
hasError=`protoc --version| grep "error"`
if [ "${hasError}" != "" ]; then
  cat > /etc/ld.so.conf.d/libprotobuf.conf << EOF
/usr/local/lib
EOF
sudo ldconfig
fi
hasError=`protoc --version| grep "error"`
if [ "${hasError}" == "" ]; then
  go get github.com/golang/protobuf
  go install github.com/golang/protobuf/protoc-gen-go/
  protoc --go_out=. hello.proto
else
  echo "has wrong"
fi