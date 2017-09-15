package effects_test

import (
	"github.com/nestorsalceda/maverik/devices"
	"github.com/nestorsalceda/maverik/effects"
	"github.com/nestorsalceda/maverik/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/stretchr/testify/mock"
)

var _ = Describe("Solid", func() {
	It("sends color to device", func() {
		const red = "FF0000"
		ledStrip := new(mocks.LedStrip)
		ledStrip.On("Write", Anything).Return(0, nil)
		ledStrip.On("AmountOfLeds").Return(144)

		effects.Solid(ledStrip, red)

		ledStrip.AssertCalled(GinkgoT(), "Write", redStripeFor(ledStrip))
	})
})

func redStripeFor(ledStrip devices.LedStrip) []byte {
	buf := make([]byte, ledStrip.AmountOfLeds()*3)
	for i := 0; i < len(buf); i += 3 {
		buf[i] = 0xff
		buf[i+1] = 0
		buf[i+2] = 0
	}

	return buf
}
