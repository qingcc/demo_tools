#!/bin/sh

source_path=http.go
image_name=http_test

echo "===> building container image"

build_result="$(go build $source_path)"

if [[ $build_result =~ ":" ]] ; then
    echo "**** encounter building error, exit"
    echo "$build_result"
    exit
fi

docker rmi -f $image_name
docker build -t $image_name  .

echo '-> ** tagging http_test/'$image_name
docker tag $image_name http_test/$image_name
