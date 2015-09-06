package ta

// Stoch calculates the Stochastic Oscillator
func StochRSI(values []float64, period int, kPeriod int, dPeriod int) ([]float64,[]float64) {
  rsis := Rsi(values, period)
  return Stoch(rsis, rsis, rsis, kPeriod, dPeriod) // It's lazy to use the existing function that expect highs/lows/closes, but that's what TA-Lib does
}
