package ta

import (
  "math"
)

type HighLowClose struct {
  High float64
  Low float64
  Close float64
}

// Atr calculates the Average True Range of an array of HighLowClose structs given a period
func Atr(values []HighLowClose, period int) []float64 {
  var atrs []float64
  startIdx := period
  endIdx := len(values)-1
  tr := TrueRanges(values)
  sma := Sma(tr[:startIdx], period)
  atrs = append(atrs, sma[len(sma)-1])
  for i := startIdx; i < endIdx; i++ {
    lastAtr := atrs[len(atrs)-1]
    atr := ((lastAtr*float64(period-1)) + tr[i]) / float64(period)
    atrs = append(atrs, atr)
  }
  return atrs
}

func TrueRanges(values []HighLowClose) []float64 {
  var result []float64
  for idx,hlc := range values[1:] {
    tr := TrueRange(hlc.High, hlc.Low, values[idx].Close)
    result = append(result, tr)
  }
  return result
}

func TrueRange(curHigh float64, curLow float64, prevClose float64) float64 {
  tempMax := math.Max(curHigh-curLow, math.Abs(curHigh - prevClose))
  return math.Max(tempMax, math.Abs(curLow - prevClose))
}
