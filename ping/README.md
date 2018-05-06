# ping - a go wrapper for ping command

[Documentation](https://godoc.org/github.com/xellio/tools/ping)

### Example
```
ip := net.ParseIP("8.8.8.8")
res, err := ping.Once(ip)
if err != nil {
	fmt.Println(err)
}
fmt.Println(string(res.Raw))
```