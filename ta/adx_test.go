package ta_test

import (
	. "github.com/vertis/gota/ta"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)
var _ = Describe("Adx", func() {
	It("should return the correct values", func() {
		in := []HighLowClose {
			HighLowClose{ High: 1.1, Low: 0.9, Close: 1.0 },
			HighLowClose{ High: 2.1, Low: 1.9, Close: 2.0 },
			HighLowClose{ High: 3.1, Low: 2.9, Close: 3.0 },
			HighLowClose{ High: 4.1, Low: 3.9, Close: 4.0 },
			HighLowClose{ High: 5.2, Low: 4.9, Close: 5.0 },
			HighLowClose{ High: 6.2, Low: 5.9, Close: 6.0 },
			HighLowClose{ High: 3.1, Low: 2.9, Close: 3.0 },
		}
		out := []float64{ 100.0, 74.59670572253354 }
		result := Adx(in, 3)
		Expect(result).To(Equal(out))
	})

	It("should return the more correct values", func() {
		in := []HighLowClose {
			HighLowClose{ High: 417.27, Low: 411.04, Close: 417.26 },
			HighLowClose{ High: 419.79, Low: 416.16, Close: 419.34 },
			HighLowClose{ High: 419.44, Low: 416.92, Close: 417.96 },
			HighLowClose{ High: 417.96, Low: 415.2, Close: 417.4 },
			HighLowClose{ High: 420.23, Low: 415.02, Close: 418.1 },
			HighLowClose{ High: 420.5, Low: 415.85, Close: 417.61 },
			HighLowClose{ High: 417.62, Low: 413.31, Close: 415.1 },
			HighLowClose{ High: 415.36, Low: 413.54, Close: 414.34 },
			HighLowClose{ High: 420.44, Low: 414.32, Close: 420.44 },
			HighLowClose{ High: 421.18, Low: 418.79, Close: 420.77 },
			HighLowClose{ High: 420.85, Low: 415.37, Close: 418.21 },
			HighLowClose{ High: 419.45, Low: 416.0, Close: 418.86 },
			HighLowClose{ High: 418.86, Low: 415.8, Close: 416.36 },
			HighLowClose{ High: 416.39, Low: 411.32, Close: 412.64 },
			HighLowClose{ High: 418.13, Low: 412.49, Close: 418.13 },
			HighLowClose{ High: 419.78, Low: 414.36, Close: 414.96 },
			HighLowClose{ High: 417.27, Low: 414.29, Close: 415.48 },
			HighLowClose{ High: 416.84, Low: 414.48, Close: 414.99 },
			HighLowClose{ High: 416.41, Low: 414.54, Close: 414.96 },
			HighLowClose{ High: 417.83, Low: 409.17, Close: 410.34 },
			HighLowClose{ High: 412.17, Low: 409.26, Close: 411.62 },
			HighLowClose{ High: 412.63, Low: 408.64, Close: 408.78 },
			HighLowClose{ High: 409.95, Low: 407.45, Close: 409.53 },
			HighLowClose{ High: 413.85, Low: 409.28, Close: 413.85 },
			HighLowClose{ High: 416.17, Low: 413.18, Close: 413.84 },
			HighLowClose{ High: 414.55, Low: 411.93, Close: 413.82 },
			HighLowClose{ High: 415.29, Low: 408.04, Close: 411.09 },
			HighLowClose{ High: 413.77, Low: 411.07, Close: 413.77 },
			HighLowClose{ High: 414.38, Low: 412.24, Close: 413.76 },
		}
		out := []float64{ 11.022182600653604, 11.034848955077214 }
		result := Adx(in, 14)
		Expect(result).To(Equal(out))
	})
})
