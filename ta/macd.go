package ta

import (
  "fmt"
  "math"
)


// Macd calculates the Moving Average Convergence Divergence indicator for an
// array of float64 values given a fastPeriod, slowPeriod and signalPeriod
func Macd(values []float64, fastPeriod int, slowPeriod int, signalPeriod int) ([]float64, []float64, []float64) {
  fastPeriodValues := Ema(values, fastPeriod)
  slowPeriodValues := Ema(values, slowPeriod)
  macdValues := minusArray(fastPeriodValues, slowPeriodValues)
  signalPeriodValues := Ema(macdValues, signalPeriod)
  macdValues, signalPeriodValues = normalizeSlice(macdValues, signalPeriodValues)
  histogram := minusArray(macdValues, signalPeriodValues)
  return macdValues, signalPeriodValues, histogram
}

func minusArray(valuesA []float64, valuesB []float64) []float64 {
  sliceA, sliceB := normalizeSlice(valuesA, valuesB)
  fmt.Println(len(sliceA), len(sliceB))
  var result []float64
  for index,element := range sliceA {
    result = append(result, element-sliceB[index])
  }
  return result
}

func normalizeSlice(valuesA []float64, valuesB []float64) ([]float64, []float64) {
  offsetA := int(math.Max(0, float64(len(valuesA)-len(valuesB))))
  offsetB := int(math.Max(0, float64(len(valuesB)-len(valuesB))))
  sliceA := valuesA[offsetA:]
  sliceB := valuesB[offsetB:]
  return sliceA, sliceB
}
