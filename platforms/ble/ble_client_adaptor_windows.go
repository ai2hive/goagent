package ble

import (
	"errors"

	blelib "github.com/hybridgroup/ble"
)

func defaultDevice(impl string) (d blelib.Device, err error) {
	return nil, errors.New("Not yet implemented for this OS.")
}
