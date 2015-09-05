package ta

import (
  "math"
)

func BollingerBands(values []float64, period int) ([]float64, []float64, []float64) {
  deviationsUp := 2.0 // TODO: change so it's possible to configure this
  deviationsDown := 2.0

  middleBand := Sma(values, period)
  offset := len(values)-len(middleBand)
  var upperBand []float64
  var lowerBand []float64
  for idx, v := range middleBand {
    backIdx := offset+idx-period+1
    curIdx := offset+idx+1
    if backIdx < 0 {
      backIdx = 0
    }
    stdDev := SliceStdDev(values[backIdx:curIdx])
    upperBand = append(upperBand, v + (stdDev * deviationsUp))
    lowerBand = append(lowerBand, v - (stdDev * deviationsDown))
  }
  return upperBand, middleBand, lowerBand
}

// SliceVariance returns the variance of the slice of float64.
func SliceVariance(values []float64) float64 {
	if 0 == len(values) {
		return 0.0
	}
	m := Mean(values)
	var sum float64
	for _, v := range values {
		d := v - m
		sum += d * d
	}
	return sum / float64(len(values))
}

// SliceStdDev returns the standard deviation of the slice of float64.
func SliceStdDev(values []float64) float64 {
	return math.Sqrt(SliceVariance(values))
}
