package ta_test

import (
	. "github.com/vertis/gota/ta"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// NB: The values here are minutely different from the ones TA-LIB returns for the same *in* values. I am unsure why, but given it is in the 14th decimal place, I've decided not to waste time troubleshooting
var _ = Describe("BollingerBands", func() {
	It("should return the correct values", func() {
		in := []float64{ 26.0, 54.0, 8.0, 77.0, 61.0, 39.0, 44.0, 91.0, 98.0, 17.0 }
		upper_band_out := []float64{ 94.92564730599291, 94.50588827974477, 92.12752961253167, 101.58367006802706, 114.64331379078675, 120.58088881180322 }
		middle_band_out := []float64{ 45.2, 47.8, 45.8, 62.4, 66.6, 57.8 }
		lower_band_out := []float64{ -4.525647305992905, 1.0941117202552277, -0.52752961253168, 23.216329931972936, 18.556686209213247, -4.980888811803233 }
		upper_band, middle_band, lower_band := BollingerBands(in, 5)
		Expect(upper_band).To(Equal(upper_band_out))
		Expect(middle_band).To(Equal(middle_band_out))
		Expect(lower_band).To(Equal(lower_band_out))
	})
})
