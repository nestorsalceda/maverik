package effects

import (
	"log"
	"strconv"

	"github.com/nestorsalceda/maverik/devices"
)

func Solid(ledStrip devices.LedStrip, color string) {
	rgb, err := strconv.ParseUint(color, 16, 32)
	if err != nil {
		log.Fatal(err)
	}

	r := byte(rgb >> 16)
	g := byte(rgb >> 8)
	b := byte(rgb)

	buf := make([]byte, 144*3)
	for i := 0; i < len(buf); i += 3 {
		buf[i] = r
		buf[i+1] = g
		buf[i+2] = b
	}

	ledStrip.Write(buf)
}
