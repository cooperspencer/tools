package acpi

import "os/exec"

var path string

func init() {
	var err error
	path, err = exec.LookPath("acpi")
	if err != nil {
		panic(err)
	}
}

//
// Battery runs `acpi -b` command and returns the raw output
//
func Battery() ([]byte, error) {
	return exec.Command(path, "-b").Output()
}

//
// AcAdapter runs `acpi -a` command and returns the raw output
//
func AcAdapter() ([]byte, error) {
	return exec.Command(path, "-a").Output()
}

//
// Thermal runs `acpi -t` command and returns the raw output
//
func Thermal() ([]byte, error) {
	return exec.Command(path, "-t").Output()
}

//
// Cooling runs `acpi -c` command and returns the raw output
//
func Cooling() ([]byte, error) {
	return exec.Command(path, "-c").Output()
}

//
// Everything runs `acpi -V` command and returns the raw output
//
func Everything() ([]byte, error) {
	return exec.Command(path, "-V").Output()
}

//
// ShowEmpty runs `acpi -s` command and returns the raw output
//
func ShowEmpty() ([]byte, error) {
	return exec.Command(path, "-s").Output()
}

//
// Details runs `acpi -i` command and returns the raw output
//
func Details() ([]byte, error) {
	return exec.Command(path, "-i").Output()
}

//
// Raw runs the acpi command with the given arguments
//
func Raw(args ...string) ([]byte, error) {
	return exec.Command(path, args...).Output()
}
