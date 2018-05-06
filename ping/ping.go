package ping

import (
	"net"
	"os/exec"
	"strconv"
	"strings"
)

//
// Result struct
//
type Result struct {
	Ip        net.IP
	Raw       []byte
	Meta      *ResultMeta
	Data      []*ResultData
	Statistic *ResultStatistic
}

//
// Struct for storing meta information in Result.Meta
//
type ResultMeta struct {
	raw   string
	Host  string
	Ip    net.IP
	Bytes string
}

//
// Struct for storing statistic information in Result.Statistic
//
type ResultStatistic struct {
	raw               []string
	PacketCount       int
	ReceivedCount     int
	PacketLossPercent float64
	Time              float64
	RTT               *ResultStatisticRTT
}

//
// Struct for storing round trip times (RTT)
//
type ResultStatisticRTT struct {
	raw  []string
	Min  float64
	Avg  float64
	Max  float64
	MDev float64
}

//
// Struct for storing data information in Result.Data
//
type ResultData struct {
	raw     string
	IcmpSeq int
	Ttl     int
	Time    float64
}

//
// Returns r.Meta as string
//
func (m *ResultMeta) String() string {
	return m.raw
}

//
// Returns r.Statistic as string
//
func (s *ResultStatistic) String() string {
	return strings.Join(s.raw, "\n")
}

//
// Returns r.Statistic.RTT as string
//
func (rtt *ResultStatisticRTT) String() string {
	return strings.Join(rtt.raw, "\n")
}

//
// Returns r.Raw as string
//
func (r *Result) String() string {
	return string(r.Raw)
}

//
// Ping the given net.IP once
//
func Once(ip net.IP, args ...string) (r *Result, err error) {
	r = &Result{
		Ip:        ip,
		Meta:      &ResultMeta{},
		Statistic: &ResultStatistic{},
	}

	args = append([]string{ip.String(), "-c 1"}, args...)
	err = r.execute(args)
	return
}

//
// Executes the ping command and calls r.parseRaw func
//
func (r *Result) execute(args []string) error {
	path, err := exec.LookPath("ping")
	if err != nil {
		return err
	}

	out, err := exec.Command(path, args...).Output()
	if err != nil {
		return err
	}
	r.Raw = out

	err = r.parseResult()
	if err != nil {
		return err
	}

	return nil
}

//
// Parse the ping result
//
func (r *Result) parseResult() error {

	lines := strings.Split(string(r.Raw), "\n")

	statsBlockStart := len(lines) - 4

	for key, line := range lines {
		switch {
		case key == 0:
			err := r.parseMetaLine(line)
			if err != nil {
				return err
			}
			break
		case key >= statsBlockStart:

			err := r.parseStatisticLine(line)
			if err != nil {
				return err
			}
			break
		default:

			lineData, ok := parseDataLine(line)
			if !ok {
				continue
			}
			r.Data = append(r.Data, lineData)
		}
	}

	return nil
}

//
// Parse line for ResultData
//
func parseDataLine(line string) (lineData *ResultData, ok bool) {

	lineData = &ResultData{
		raw: line,
	}

	splittedDataLine := strings.Split(line, ": ")
	if len(splittedDataLine) < 2 {
		return
	}

	splittedData := strings.Split(splittedDataLine[1], " ")

	for _, pair := range splittedData {
		kV := strings.Split(pair, "=")
		switch kV[0] {
		case "icmp_seq":
			val, err := strconv.Atoi(kV[1])
			if err != nil {
				return
			}
			lineData.IcmpSeq = val
			break
		case "ttl":
			val, err := strconv.Atoi(kV[1])
			if err != nil {
				return
			}
			lineData.Ttl = val
			break
		case "time":
			val, err := strconv.ParseFloat(kV[1], 32)
			if err != nil {
				return
			}
			lineData.Time = val
			break
		}
	}
	ok = true
	return

}

//
// Parse line for statistic
//
func (r *Result) parseStatisticLine(line string) error {
	r.Statistic.raw = append(r.Statistic.raw, line)
	return nil
}

//
// Parse line for meta
//
func (r *Result) parseMetaLine(line string) error {
	r.Meta = &ResultMeta{
		raw: line,
	}

	splitted := strings.Split(line, " ")
	r.Meta.Host = splitted[1]

	parsedIp := net.ParseIP(splitted[2][1 : len(splitted[2])-1])
	r.Meta.Ip = parsedIp

	r.Meta.Bytes = splitted[3]

	return nil
}
