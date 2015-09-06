package ta

func Roc(values []float64, period int) []float64 {
  var results []float64
  if period > len(values) {
    return results
  }
  for idx, el := range values[period:] {
    roc := ((el - values[idx]) / values[idx]) *  100
    results = append(results, roc)
  }
  return results
}
