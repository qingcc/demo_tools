module github.com/qingcc/demo_tools

go 1.13

require (
	github.com/anacrolix/sync v0.2.0 // indirect
	github.com/apache/thrift v0.13.0 // indirect
	github.com/coreos/etcd v3.3.17+incompatible
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/golang/groupcache v0.0.0-20191027212112-611e8accdfc9 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/huandu/xstrings v1.2.1 // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	github.com/klauspost/reedsolomon v1.9.3 // indirect
	github.com/lucas-clemente/quic-go v0.13.1 // indirect
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-isatty v0.0.10 // indirect
	github.com/nacos-group/nacos-sdk-go v0.0.0-20190820112454-5245ea3cded6
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/pkg/profile v1.4.0
	github.com/prometheus/client_golang v1.11.1
	github.com/rcrowley/go-metrics v0.0.0-20190826022208-cac0b30c2563
	github.com/robfig/cron v1.2.0
	github.com/smallnest/rpcx v0.0.0-20191202025149-2fd1f4f7e90c
	github.com/tealeg/xlsx v1.0.5
	github.com/xtaci/kcp-go v5.4.19+incompatible // indirect
	go.opencensus.io v0.22.2 // indirect
	go.uber.org/zap v1.12.0
	google.golang.org/appengine v1.6.5 // indirect
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.26.0-rc.1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

replace (
	github.com/qingcc/demo_tools/util => ../demo_tools/util
	golang.org/x/crypto v0.0.0-20180820150726-614d502a4dac => github.com/golang/crypto v0.0.0-20180820150726-614d502a4dac
	golang.org/x/net v0.0.0-20180821023952-922f4815f713 => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
)
