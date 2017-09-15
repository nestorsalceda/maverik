package devices

type LedStrip interface {
	Write(leds []byte) (int, error)
}
