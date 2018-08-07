# ping - a go wrapper for ping command

[![go report card](https://goreportcard.com/badge/github.com/xellio/tools "go report card")](https://goreportcard.com/report/github.com/xellio/tools)
[![MIT license](http://img.shields.io/badge/license-MIT-brightgreen.svg)](http://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/xellio/tools/ping?status.svg)](https://godoc.org/github.com/xellio/tools/ping)

### Example
```
ip := net.ParseIP("8.8.8.8")
res, err := ping.Once(ip)
if err != nil {
	fmt.Println(err)
}
fmt.Println(string(res.Raw))
```