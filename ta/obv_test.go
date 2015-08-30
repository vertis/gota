package ta_test

import (
	. "github.com/vertis/gota/ta"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)
//[53.591160220994475, 55.140186915887845, 67.77751138815631, 69.38322227917307, 40.32069898478417, nil, nil, nil, nil, nil]
var _ = Describe("Obv", func() {
	It("should return the correct values", func() {
		in := []CloseVolume{ CloseVolume { 26.0, 100.0 }, CloseVolume { 54.0, 100.0 }, CloseVolume { 55.0, 100.0 }, CloseVolume { 54.0, 100.0 } }
		out := []float64{ 100.0, 200.0, 100.0 }
		result := Obv(in)
		Expect(result).To(Equal(out))
	})
})
