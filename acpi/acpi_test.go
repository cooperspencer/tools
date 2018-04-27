// +build !testing

package acpi

import (
	"testing"

	"github.com/stretchr/testify/require"
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
	parsed, err := Parse(getData())
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

func BenchmarkParse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = Parse(getData())
	}
}
