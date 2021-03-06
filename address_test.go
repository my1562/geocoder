package geocoder

import (
	"testing"

	"github.com/bradleyjkemp/cupaloy"
)

func TestAddress_GetBuildingAsString(t *testing.T) {
	geo := NewGeocoder()

	buildings := []string{
		geo.AddressByID(100003).Address.GetBuildingAsString(),
		geo.AddressByID(108527).Address.GetBuildingAsString(),
		geo.AddressByID(113134).Address.GetBuildingAsString(),
	}
	cupaloy.SnapshotT(t, buildings)
}
