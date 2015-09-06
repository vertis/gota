package ta

import (
  "math"
)

// Stoch calculates the Stochastic Oscillator
func Stoch(highs []float64, lows []float64, closes []float64, period int, smaPeriod int) ([]float64,[]float64) {
  // TODO: validate that all the arrays are the same length
  var stochs []float64
  for i := period-1; i < len(highs); i++ {
    highestHigh := max(highs[i-period+1:i+1])
    lowestLow := min(lows[i-period+1:i+1])
    stoch := (closes[i] - lowestLow)/(highestHigh - lowestLow) * 100
    stochs = append(stochs, stoch)
  }
  stochSignals := Sma(stochs, smaPeriod)
  return stochs, stochSignals
}

func max(values []float64) float64 {
  curMax := values[0]
  for _,el := range values[1:] {
    curMax = math.Max(curMax, el)
  }
  return curMax
}

func min(values []float64) float64 {
  curMin := values[0]
  for _,el := range values[1:] {
    curMin = math.Min(curMin, el)
  }
  return curMin
}
