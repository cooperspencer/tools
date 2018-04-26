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

	require.Equal(t, parsed.BatteryInformation[0].Number, 0)
	require.Equal(t, parsed.BatteryInformation[0].Level, 100)
	require.Equal(t, parsed.BatteryInformation[0].Status, "Full")
	require.Equal(t, parsed.BatteryInformation[0].DesignCapacity, 3464)
	require.Equal(t, parsed.BatteryInformation[0].LastFullCapacity, 2864)
	require.Equal(t, parsed.BatteryInformation[0].LastFullCapacityPercent, 82)

	require.Equal(t, parsed.BatteryInformation[1].Number, 1)
	require.Equal(t, parsed.BatteryInformation[1].Level, 96)
	require.Equal(t, parsed.BatteryInformation[1].Status, "Unknown")
	require.Equal(t, parsed.BatteryInformation[1].DesignCapacity, 1877)
	require.Equal(t, parsed.BatteryInformation[1].LastFullCapacity, 1446)
	require.Equal(t, parsed.BatteryInformation[1].LastFullCapacityPercent, 77)

	require.NotEmpty(t, parsed.AdapterInformation)
	require.Len(t, parsed.AdapterInformation, 1)

	require.Equal(t, parsed.AdapterInformation[0].Number, 0)
	require.Equal(t, parsed.AdapterInformation[0].Status, "on-line")

	require.NotEmpty(t, parsed.CoolingInformation)
	require.Len(t, parsed.CoolingInformation, 7)

	require.Equal(t, parsed.CoolingInformation[0].Number, 0)
	require.Equal(t, parsed.CoolingInformation[0].Processor, 0)
	require.Equal(t, parsed.CoolingInformation[0].ProcessorMax, 10)
	require.Equal(t, parsed.CoolingInformation[0].Note, "")

	require.Equal(t, parsed.CoolingInformation[1].Number, 1)
	require.Equal(t, parsed.CoolingInformation[1].Processor, 0)
	require.Equal(t, parsed.CoolingInformation[1].ProcessorMax, 10)
	require.Equal(t, parsed.CoolingInformation[1].Note, "")

	require.Equal(t, parsed.CoolingInformation[2].Number, 2)
	require.Equal(t, parsed.CoolingInformation[2].Processor, 0)
	require.Equal(t, parsed.CoolingInformation[2].ProcessorMax, 0)
	require.Equal(t, parsed.CoolingInformation[2].Note, "intel_powerclamp no state information available")

	require.Equal(t, parsed.CoolingInformation[3].Number, 3)
	require.Equal(t, parsed.CoolingInformation[3].Processor, 0)
	require.Equal(t, parsed.CoolingInformation[3].ProcessorMax, 0)
	require.Equal(t, parsed.CoolingInformation[3].Note, "x86_pkg_temp no state information available")

	require.Equal(t, parsed.CoolingInformation[4].Number, 4)
	require.Equal(t, parsed.CoolingInformation[4].Processor, 0)
	require.Equal(t, parsed.CoolingInformation[4].ProcessorMax, 10)
	require.Equal(t, parsed.CoolingInformation[4].Note, "")

	require.Equal(t, parsed.CoolingInformation[5].Number, 5)
	require.Equal(t, parsed.CoolingInformation[5].Processor, 0)
	require.Equal(t, parsed.CoolingInformation[5].ProcessorMax, 10)
	require.Equal(t, parsed.CoolingInformation[5].Note, "")

	require.Equal(t, parsed.CoolingInformation[6].Number, 6)
	require.Equal(t, parsed.CoolingInformation[6].Processor, 0)
	require.Equal(t, parsed.CoolingInformation[6].ProcessorMax, 0)
	require.Equal(t, parsed.CoolingInformation[6].Note, "pch_wildcat_point no state information available")

	require.NotEmpty(t, parsed.ThermalInformation)
	require.Len(t, parsed.ThermalInformation, 1)
}

func BenchmarkParse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = Parse(getData())
	}
}
