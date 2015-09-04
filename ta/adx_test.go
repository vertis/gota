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
})
