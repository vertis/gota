package ta_test

import (
	. "github.com/vertis/gota/ta"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)


var _ = Describe("Roc", func() {
	It("should return the correct values", func() {
		in := []float64 { 26.0, 54.0, 8.0, 77.0, 61.0, 39.0, 44.0, 91.0, 98.0, 17.0 }
		out := []float64{ 50.0, -18.51851851851852, 1037.5, 27.27272727272727, -72.1311475409836 }
		Expect(Roc(in, 5)).To(Equal(out))
	})
})
