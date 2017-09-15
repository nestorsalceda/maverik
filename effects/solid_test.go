package effects_test

import (
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

		effects.Solid(ledStrip, red)

		ledStrip.AssertCalled(GinkgoT(), "Write", redStripe())
	})
})

func redStripe() []byte {
	buf := make([]byte, 144*3)
	for i := 0; i < len(buf); i += 3 {
		buf[i] = 0xff
		buf[i+1] = 0
		buf[i+2] = 0
	}

	return buf
}
