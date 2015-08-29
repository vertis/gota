package ta_test

import (
	. "github.com/vertis/gota/ta"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Aroon", func() {
	It("should return the correct values", func() {
		in := []float64{ 26.0, 54.0, 8.0, 77.0, 61.0, 39.0, 44.0, 91.0, 98.0, 17.0 }
		aroon_down := []float64{40.0, 20.0, 0.0, 40.0, 100.0}
		aroon_up := []float64{60.0, 40.0, 100.0, 100.0, 80.0}
		res_aroon_down, res_aroon_up := Aroon(in, 5)
		Expect(res_aroon_down).To(Equal(aroon_down))
		Expect(res_aroon_up).To(Equal(aroon_up))
	})
})
