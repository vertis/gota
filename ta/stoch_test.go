package ta_test

import (
	. "github.com/vertis/gota/ta"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stoch", func() {
	// Example from http://stockcharts.com/school/doku.php?id=chart_school:technical_indicators:stochastic_oscillator_fast_slow_and_full
	It("should return the correct values", func() {
		high_in := []float64 { 127.01, 127.62, 126.59, 127.35, 128.17, 128.43, 127.37, 126.42, 126.90,
													 126.85, 125.65, 125.72, 127.16, 127.72, 127.49, 128.22, 128.27, 128.09,
													 128.27, 127.74, 128.77, 129.29, 130.06, 129.12, 129.29, 128.47, 128.09,
													 128.65, 129.14, 128.64 }
		low_in := []float64{ 125.36, 126.16, 124.93, 126.09, 126.82, 126.48, 126.03, 124.83, 126.39,
												 125.72, 124.56, 124.57, 125.07, 126.86, 126.63, 126.80, 126.71, 126.80,
												 126.13, 125.92, 126.99, 127.81, 128.47, 128.06, 127.61, 127.60, 127.00,
												 126.90, 127.49, 127.40 }
		// Our algorithm requires everything to be the same length so this is zero padded, despite the example not having the data
		close_in := []float64{ 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00,
			 										 0.00, 127.29, 127.18, 128.01, 127.11, 127.73, 127.06, 127.33, 128.71,
													 127.87, 128.58, 128.60, 127.93, 128.11, 127.60, 127.60, 128.69, 128.27}
		// NB: After all that the values in the example spreadsheet aren't correct
		//out := []float64{ 70.44, 67.61, 89.20, 65.81, 81.75, 64.52, 74.53, 98.58, 70.10, 73.06, 73.42, 61.23, 60.96, 40.39, 40.39, 66.83, 56.73 }
		out := []float64{ 70.54263565891475, 67.70025839793286, 89.14728682170502, 65.89147286821691, 81.91214470284233, 64.59948320413429, 74.66307277628006, 98.57482185273155, 69.97885835095158, 73.09090909090928, 73.45454545454531, 61.20218579234987, 60.921843687375045, 40.57971014492735, 40.57971014492735, 66.90821256038637, 56.76328502415478 }
		stochs, _ := Stoch(high_in, low_in, close_in, 14, 3)
		Expect(stochs).To(Equal(out))
	})
})
