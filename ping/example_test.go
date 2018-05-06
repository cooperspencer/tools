package ping_test

import (
	"fmt"
	"net"

	"github.com/xellio/tools/ping"
)

//
// Example of Once function
//
func ExampleOnce() {
	ip := net.ParseIP("8.8.8.8")
	res, err := ping.Once(ip)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Meta.String())

	//Output:
	//PING 8.8.8.8 (8.8.8.8) 56(84) bytes of data.
}
