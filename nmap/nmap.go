package nmap

import (
	"os/exec"

	"github.com/xellio/ns"
)

var path string

func init() {
	var err error
	path, err = exec.LookPath("nmap")
	if err != nil {
		panic(err)
	}
}

//
// Query runs the nmap command with the given arguments
//
func Query(args ...string) (ns.Result, error) {
	return ns.Query(args)
}
