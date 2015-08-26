package ta

import (
  "fmt"
  "math"
)

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
