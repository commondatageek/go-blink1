package blink1

/*
#cgo CFLAGS: -I/opt/homebrew/Cellar/libusb-compat/0.1.8/include -I/opt/homebrew/Cellar/libusb/1.0.27/include
#cgo LDFLAGS: /opt/homebrew/Cellar/libusb-compat/0.1.8/lib/libusb.a /opt/homebrew/Cellar/libusb/1.0.27/lib/libusb-1.0.a -framework CoreFoundation -framework IOKit
#include <usb.h>
*/
import "C"

import (
	"time"

	"github.com/commondatageek/go-blink1/libusb"
)

func fadeToRgbBlink1(device *Device, fadeTime time.Duration, red, green, blue, led uint8) (bytesWritten int) {
	bytesWritten = libusb.SendBlink1Command(device.Device, toMs(fadeTime), red, blue, green, led)
	return
}
