package acpi

import (
	"bytes"
	"log"
	"os/exec"
	"strconv"
)

// path stores information on the acpi location
var path string

//
// ACPI stores the parsed output of acpi
//
type ACPI struct {
	BatteryInformation []*BatteryInformation
	AdapterInformation []*AdapterInformation
	ThermalInformation []*ThermalInformation
	CoolingInformation []*CoolingInformation
}

//
// BatteryInformation stores parsed information for a battery
//
type BatteryInformation struct {
	Number                  int
	Status                  string
	Level                   int
	DesignCapacity          int
	LastFullCapacity        int
	LastFullCapacityPercent int
}

//
// AdapterInformation stores parsed information for an adapter
//
type AdapterInformation struct {
	Number int
	Status string
}

//
// ThermalInformation stores parsed thermal information
//
type ThermalInformation struct {
	Number            int
	Status            string
	Degree            float32
	Unit              string
	CriticalTripPoint float32
}

//
// CoolingInformation stores parsed information for cooling
//
type CoolingInformation struct {
	Number       int
	Processor    int
	ProcessorMax int
	Note         string
	// add state information
}

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

//
// Parse converts the given raw acpi output to an ACPI struct
//
func Parse(raw []byte) (*ACPI, error) {
	lines := bytes.Split(raw[:len(raw)-1], []byte("\n"))

	info := &ACPI{}

	var tmpMap = make(map[string]map[int][][]byte)

	for _, l := range lines {
		kv := bytes.Split(l, []byte(":"))

		if len(tmpMap[string(kv[0][:7])]) <= 0 {
			tmpMap[string(kv[0][:7])] = make(map[int][][]byte)
		}

		number, err := strconv.Atoi(string(bytes.TrimSpace(kv[0][7:])))
		if err == nil {
			tmpMap[string(kv[0][:7])][number] = append(tmpMap[string(kv[0][:7])][number], kv[1])
		}
	}

	for k, v := range tmpMap {
		switch k {
		case "Battery":
			batteryInformation, err := parseBattery(v)
			if err != nil {
				continue
			}
			info.BatteryInformation = batteryInformation
			break
		case "Adapter":
			adapterInformation, err := parseAdapter(v)
			if err != nil {
				continue
			}
			info.AdapterInformation = adapterInformation
			break
		case "Thermal":
			thermalInformation, err := parseThermal(v)
			if err != nil {
				continue
			}
			info.ThermalInformation = thermalInformation
			break
		case "Cooling":
			coolingInformation, err := parseCooling(v)
			if err != nil {
				continue
			}
			info.CoolingInformation = coolingInformation
			break
		default:
			log.Println("unknown key " + k)
			break
		}
	}
	return info, nil
}

//
// parseBattery ...
//
func parseBattery(raw map[int][][]byte) (information []*BatteryInformation, err error) {
	for key, values := range raw {
		info := &BatteryInformation{}
		info.Number = key
		for _, value := range values {

			splitted := bytes.Split(value, []byte(","))

			if bytes.Contains(splitted[0], []byte("capacity")) {
				designSplit := bytes.Split(bytes.TrimSpace(splitted[0]), []byte("design capacity"))

				dc, err := strconv.Atoi(string(bytes.TrimSpace(designSplit[1][:len(designSplit[1])-3])))
				if err == nil {
					info.DesignCapacity = dc
				}

				lastFullSplit := bytes.Split(bytes.TrimSpace(splitted[1]), []byte("last full capacity"))

				lastFullValues := bytes.Split(lastFullSplit[1], []byte("="))
				lastFullValues[0] = bytes.TrimSpace(lastFullValues[0])
				lastFullValues[1] = bytes.TrimSpace(lastFullValues[1])

				lfc, err := strconv.Atoi(string(bytes.TrimSpace(lastFullValues[0][:len(lastFullValues[0])-3])))
				if err == nil {
					info.LastFullCapacity = lfc
				}

				lfcp, err := strconv.Atoi(string(bytes.TrimSpace(lastFullValues[1][:len(lastFullValues[1])-1])))
				if err == nil {
					info.LastFullCapacityPercent = lfcp
				}

				continue
			}

			// non capacity line
			info.Status = string(bytes.TrimSpace(splitted[0]))
			level, err := strconv.Atoi(string(bytes.TrimSpace(splitted[1][:len(splitted[1])-1])))
			if err == nil {
				info.Level = level
			}
		}

		information = append(information, info)
	}
	return
}

//
// parseAdapter ..
//
func parseAdapter(raw map[int][][]byte) (information []*AdapterInformation, err error) {
	for key, values := range raw {
		info := &AdapterInformation{}
		info.Number = key
		for _, value := range values {
			info.Status = string(bytes.TrimSpace(value))
		}
		information = append(information, info)
	}
	return
}

//
// parseThermal ...
//
func parseThermal(raw map[int][][]byte) (information []*ThermalInformation, err error) {
	for key, values := range raw {
		info := &ThermalInformation{}
		info.Number = key
		for _, value := range values {
			splittedVal := bytes.Split(value, []byte(","))
			//todo
			if len(splittedVal) == 2 {
				//info.Status
				//info.Degree
				//info.Unit
				continue
			}
			//info.CriticalTripPoint
		}
		information = append(information, info)
	}
	return
}

//
// parseCooling ...
//
func parseCooling(raw map[int][][]byte) (information []*CoolingInformation, err error) {
	for key, values := range raw {
		info := &CoolingInformation{}
		info.Number = key
		for _, value := range values {

			if bytes.Equal(bytes.TrimSpace(value)[0:9], []byte("Processor")) {
				splittedProc := bytes.Split(bytes.TrimSpace(value)[9:], []byte(" of "))
				proc, err := strconv.Atoi(string(bytes.TrimSpace(splittedProc[0])))
				if err == nil {
					info.Processor = proc
				}
				procMax, err := strconv.Atoi(string(bytes.TrimSpace(splittedProc[1])))
				if err == nil {
					info.ProcessorMax = procMax
				}
				continue
			}
			info.Note = string(bytes.TrimSpace(value))
		}
		information = append(information, info)
	}
	return
}
