// +build !testing

package acpi

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func getData() []byte {
	return []byte(`Battery 0: Full, 100%
Battery 0: design capacity 3464 mAh, last full capacity 2864 mAh = 82%
Battery 1: Unknown, 96%
Battery 1: design capacity 1877 mAh, last full capacity 1446 mAh = 77%
Adapter 0: on-line
Thermal 0: ok, 42.0 degrees C
Thermal 0: trip point 0 switches to mode critical at temperature 103.0 degrees C
Cooling 0: Processor 0 of 10
Cooling 1: Processor 0 of 10
Cooling 2: intel_powerclamp no state information available
Cooling 3: x86_pkg_temp no state information available
Cooling 4: Processor 0 of 10
Cooling 5: Processor 0 of 10
Cooling 6: pch_wildcat_point no state information available
`)
}

func TestParse(t *testing.T) {
	parsed, err := parse(getData())
	require.Nil(t, err)

	require.NotEmpty(t, parsed.BatteryInformation)
	require.Len(t, parsed.BatteryInformation, 2)
	for _, battery := range parsed.BatteryInformation {
		switch battery.Number {
		case 0:
			require.Equal(t, battery.Number, 0)
			require.Equal(t, battery.Level, 100)
			require.Equal(t, battery.Status, "Full")
			require.Equal(t, battery.DesignCapacity, 3464)
			require.Equal(t, battery.LastFullCapacity, 2864)
			require.Equal(t, battery.LastFullCapacityPercent, 82)
			break
		case 1:
			require.Equal(t, battery.Number, 1)
			require.Equal(t, battery.Level, 96)
			require.Equal(t, battery.Status, "Unknown")
			require.Equal(t, battery.DesignCapacity, 1877)
			require.Equal(t, battery.LastFullCapacity, 1446)
			require.Equal(t, battery.LastFullCapacityPercent, 77)
			break
		}
	}

	require.NotEmpty(t, parsed.AdapterInformation)
	require.Len(t, parsed.AdapterInformation, 1)

	require.Equal(t, parsed.AdapterInformation[0].Number, 0)
	require.Equal(t, parsed.AdapterInformation[0].Status, "on-line")

	require.NotEmpty(t, parsed.CoolingInformation)
	require.Len(t, parsed.CoolingInformation, 7)
	for _, cooling := range parsed.CoolingInformation {
		switch cooling.Number {
		case 0:
			require.Equal(t, cooling.Number, 0)
			require.Equal(t, cooling.Processor, 0)
			require.Equal(t, cooling.ProcessorMax, 10)
			require.Equal(t, cooling.Note, "")
			break
		case 1:
			require.Equal(t, cooling.Number, 1)
			require.Equal(t, cooling.Processor, 0)
			require.Equal(t, cooling.ProcessorMax, 10)
			require.Equal(t, cooling.Note, "")
			break
		case 2:
			require.Equal(t, cooling.Number, 2)
			require.Equal(t, cooling.Processor, 0)
			require.Equal(t, cooling.ProcessorMax, 0)
			require.Equal(t, cooling.Note, "intel_powerclamp no state information available")
			break
		case 3:
			require.Equal(t, cooling.Number, 3)
			require.Equal(t, cooling.Processor, 0)
			require.Equal(t, cooling.ProcessorMax, 0)
			require.Equal(t, cooling.Note, "x86_pkg_temp no state information available")
			break
		case 4:
			require.Equal(t, cooling.Number, 4)
			require.Equal(t, cooling.Processor, 0)
			require.Equal(t, cooling.ProcessorMax, 10)
			require.Equal(t, cooling.Note, "")
			break
		case 5:
			require.Equal(t, cooling.Number, 5)
			require.Equal(t, cooling.Processor, 0)
			require.Equal(t, cooling.ProcessorMax, 10)
			require.Equal(t, cooling.Note, "")
			break
		case 6:
			require.Equal(t, cooling.Number, 6)
			require.Equal(t, cooling.Processor, 0)
			require.Equal(t, cooling.ProcessorMax, 0)
			require.Equal(t, cooling.Note, "pch_wildcat_point no state information available")
			break
		}
	}

	require.NotEmpty(t, parsed.ThermalInformation)
	require.Len(t, parsed.ThermalInformation, 1)
}

func TestBattery(t *testing.T) {
	_, err := Battery()
	require.Nil(t, err)
}

func TestParseBatteryInformation(t *testing.T) {
	raw := []byte(`Battery 0: Full, 100%
Battery 0: design capacity 3464 mAh, last full capacity 2864 mAh = 82%
Battery 1: Unknown, 96%
Battery 1: design capacity 1877 mAh, last full capacity 1446 mAh = 77%
`)
	parsed, err := parse(raw)
	require.Nil(t, err)

	bi := parsed.BatteryInformation
	require.Len(t, bi, 2)
}

func TestAcAdapter(t *testing.T) {
	_, err := AcAdapter()
	require.Nil(t, err)
}

func TestParseAcAdapter(t *testing.T) {
	raw := []byte(`Adapter 0: on-line
`)
	parsed, err := parse(raw)
	require.Nil(t, err)

	ai := parsed.AdapterInformation
	require.Len(t, ai, 1)
}

func TestThermal(t *testing.T) {
	_, err := Thermal()
	require.Nil(t, err)
}

func TestParseThermal(t *testing.T) {
	raw := []byte(`Thermal 0: ok, 42.0 degrees C
Thermal 0: trip point 0 switches to mode critical at temperature 103.0 degrees C
`)
	parsed, err := parse(raw)
	require.Nil(t, err)

	ti := parsed.ThermalInformation
	require.Len(t, ti, 1)
}

func TestCooling(t *testing.T) {
	_, err := Cooling()
	require.Nil(t, err)
}
func TestParseCooling(t *testing.T) {
	raw := []byte(`Cooling 0: Processor 0 of 10
Cooling 1: Processor 0 of 10
Cooling 2: intel_powerclamp no state information available
Cooling 3: x86_pkg_temp no state information available
Cooling 4: Processor 0 of 10
Cooling 5: Processor 0 of 10
Cooling 6: pch_wildcat_point no state information available
`)
	parsed, err := parse(raw)
	require.Nil(t, err)

	ci := parsed.CoolingInformation
	require.Len(t, ci, 7)
}

func TestRaw(t *testing.T) {
	_, err := Raw("-s", "-i", "-b")
	require.Nil(t, err)
}

func BenchmarkParse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = parse(getData())
	}
}
