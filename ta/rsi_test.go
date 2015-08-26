package ta_test

import (
	. "github.com/vertis/gota/ta"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)
//[53.591160220994475, 55.140186915887845, 67.77751138815631, 69.38322227917307, 40.32069898478417, nil, nil, nil, nil, nil]
var _ = Describe("Rsi", func() {
	It("should return the correct values", func() {
		in := []float64{ 26.0, 54.0, 8.0, 77.0, 61.0, 39.0, 44.0, 91.0, 98.0, 17.0 }
		out := []float64{ 53.591160220994475, 55.140186915887845, 67.77751138815631, 69.38322227917307, 40.32069898478417 }
		result := Rsi(in, 5)
		Expect(result).To(Equal(out))
	})
})
