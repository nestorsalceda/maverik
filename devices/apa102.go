package devices

import (
	"log"

	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/devices/apa102"
	"periph.io/x/periph/host"
)

const numberOfLeds = 144
const intensity = 100
const temperature = 5000

type APA102 struct {
	numberOfLeds int
	device       *apa102.Dev
}

func NewAPA102() LedStrip {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	s, err := spireg.Open("")
	if err != nil {
		log.Fatal(err)
	}
	s.LimitSpeed(8000000)

	device, err := apa102.New(s, numberOfLeds, intensity, temperature)
	if err != nil {
		log.Fatal(err)
	}

	return &APA102{
		numberOfLeds,
		device,
	}
}

func (d *APA102) Write(leds []byte) (int, error) {
	return d.device.Write(leds)
}
