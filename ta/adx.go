package ta

import (
  "math"
)

// Adx calculates the Average directional movement index of an array of HighLowClose structs given a period
func Adx(values []HighLowClose, period int) []float64 {
  atrs := Atr(values, period)
  var adxs []float64
  var dmPluses []float64
  var dmMinuses []float64
  var dxs []float64
  for idx,el := range values[period:] {
    offsetIdx := idx+period
    dmPlus := 0.0
    dmMinus := 0.0
    upMove := el.High - values[offsetIdx].High
    downMove := el.Low - values[offsetIdx].Low
    if upMove > downMove && upMove > 0.0 {
      dmPlus = upMove
    }
    if downMove > upMove && downMove > 0.0 {
      dmMinus = downMove
    }
    dmPluses = append(dmPluses, dmPlus)
    dmMinuses = append(dmMinuses, dmMinus)
    diPlus := 100 * (Ema(dmPluses, period)[len(dmPluses)-1] / atrs[idx])
    diMinus := 100 * (Ema(dmMinuses, period)[len(dmMinuses)-1] / atrs[idx])
    dxs = append(dxs, (math.Abs((diPlus - diMinus) / (diPlus + diMinus))))
    adx := 100 * Ema(dxs, period)[len(dxs)-1]
    adxs = append(adxs, adx)
  }
  return adxs
}
