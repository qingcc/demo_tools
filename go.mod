module github.com/qingcc/demo_tools

go 1.13

require (
	github.com/Shopify/goreferrer v0.0.0-20181106222321-ec9c9a553398 // indirect
	github.com/anacrolix/sync v0.2.0 // indirect
	github.com/apache/thrift v0.13.0 // indirect
	github.com/coreos/etcd v3.3.17+incompatible
	github.com/fatih/structs v1.1.0 // indirect
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/golang/groupcache v0.0.0-20191027212112-611e8accdfc9 // indirect
	github.com/golang/protobuf v1.4.0
	github.com/gorilla/schema v1.1.0 // indirect
	github.com/huandu/xstrings v1.2.1 // indirect
	github.com/iris-contrib/blackfriday v2.0.0+incompatible // indirect
	github.com/iris-contrib/formBinder v5.0.0+incompatible // indirect
	github.com/iris-contrib/go.uuid v2.0.0+incompatible // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	github.com/kataras/golog v0.0.11 // indirect
	github.com/kataras/iris v11.1.1+incompatible
	github.com/klauspost/compress v1.10.5 // indirect
	github.com/klauspost/reedsolomon v1.9.3 // indirect
	github.com/lucas-clemente/quic-go v0.13.1 // indirect
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-isatty v0.0.10 // indirect
	github.com/microcosm-cc/bluemonday v1.0.2 // indirect
	github.com/nacos-group/nacos-sdk-go v0.0.0-20190820112454-5245ea3cded6
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/pkg/profile v1.4.0
	github.com/prometheus/client_golang v1.1.0
	github.com/rcrowley/go-metrics v0.0.0-20190826022208-cac0b30c2563
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/smallnest/rpcx v0.0.0-20191202025149-2fd1f4f7e90c
	github.com/tealeg/xlsx v1.0.5
	github.com/xtaci/kcp-go v5.4.19+incompatible // indirect
	go.opencensus.io v0.22.2 // indirect
	go.uber.org/zap v1.12.0
	golang.org/x/crypto v0.0.0-20191128160524-b544559bb6d1 // indirect
	golang.org/x/net v0.0.0-20191126235420-ef20fe5d7933 // indirect
	google.golang.org/appengine v1.6.5 // indirect
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.21.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

replace (
	github.com/qingcc/demo_tools/util => ../demo_tools/util
	golang.org/x/crypto v0.0.0-20180820150726-614d502a4dac => github.com/golang/crypto v0.0.0-20180820150726-614d502a4dac
	golang.org/x/net v0.0.0-20180821023952-922f4815f713 => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
)
