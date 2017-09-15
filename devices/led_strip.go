package devices

type LedStrip interface {
	AmountOfLeds() int
	Write(leds []byte) (int, error)
}
