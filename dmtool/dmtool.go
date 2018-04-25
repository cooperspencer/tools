package dmtool

import (
	"os/exec"
)

var path string

func init() {
	var err error
	path, err = exec.LookPath("dm-tool")
	if err != nil {
		panic(err)
	}
}

//
// ListSeats runs `dm-tool list-seats` command and returns the raw output
//
func ListSeats() ([]byte, error) {
	return exec.Command(path, "list-seats").Output()
}

//
// Raw runs the dm-tool command with the given arguments
//
func Raw(args ...string) ([]byte, error) {
	return exec.Command(path, args...).Output()
}
