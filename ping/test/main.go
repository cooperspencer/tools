package main

import (
	"fmt"
	"net"

	"github.com/xellio/tools/ping"
)

func main() {
	ip := net.ParseIP("8.8.8.8")
	res, err := ping.Once(ip)
	if err != nil {
		fmt.Println(err)
	}

	//	fmt.Println(res.Meta.Host)
	//	fmt.Println(res.Meta.Ip)
	//	fmt.Println(res.Meta.Bytes)

	fmt.Println(res.Statistic.RTT.String())

	for _, d := range res.Data {
		fmt.Println(d.Time)
	}

}
