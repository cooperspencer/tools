// +build !testing

package acpi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	raw := []byte(`Battery 0: Full, 100%
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
Cooling 6: pch_wildcat_point no state information available`)

	parsed, err := Parse(raw)
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

}
