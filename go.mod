module github.com/qingcc/demo_tools

go 1.13

require (
	github.com/garyburd/redigo v1.6.0
	github.com/gin-gonic/gin v1.4.0
	github.com/shopspring/decimal v0.0.0-20191009025716-f1972eb1d1f5
)

replace (
	github.com/qingcc/demo_tools/util => ../demo_tools/util
	golang.org/x/crypto v0.0.0-20180820150726-614d502a4dac => github.com/golang/crypto v0.0.0-20180820150726-614d502a4dac
	golang.org/x/net v0.0.0-20180821023952-922f4815f713 => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
)
