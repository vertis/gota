package ta_test

import (
	. "github.com/vertis/gota/ta"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// { :out_real_upper_band=>[94.9256473059929, 94.50588827974477, 92.12752961253167, 101.58367006802709, 114.64331379078678, 120.58088881180323, nil, nil, nil, nil],
// :out_real_middle_band=>[45.2, 47.8, 45.8, 62.4, 66.6, 57.8, nil, nil, nil, nil],
// :out_real_lower_band=>[-4.525647305992891, 1.0941117202552206, -0.5275296125316871, 23.216329931972908, 18.55668620921321, -4.98088881180324, nil, nil, nil, nil]}
var _ = Describe("BollingerBands", func() {
	It("should return the correct values", func() {
		in := []float64{ 26.0, 54.0, 8.0, 77.0, 61.0, 39.0, 44.0, 91.0, 98.0, 17.0 }
		upper_band_out := []float64{ 94.9256473059929, 94.50588827974477, 92.12752961253167, 101.58367006802709, 114.64331379078678, 120.58088881180323}
		middle_band_out := []float64{ 45.2, 47.8, 45.8, 62.4, 66.6, 57.8 }
		lower_band_out := []float64{ -4.525647305992891, 1.0941117202552206, -0.5275296125316871, 23.216329931972908, 18.55668620921321, -4.98088881180324 }
		upper_band, middle_band, lower_band := BollingerBands(in, 5)
		Expect(upper_band).To(Equal(upper_band_out))
		Expect(middle_band).To(Equal(middle_band_out))
		Expect(lower_band).To(Equal(lower_band_out))
	})
})
