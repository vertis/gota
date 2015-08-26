package ta

import (
  
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
