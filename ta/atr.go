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
  tr := TrueRange(values)
  sma := Sma(tr[:startIdx], period)
  atrs = append(atrs, sma[len(sma)-1])
  for i := startIdx; i < endIdx; i++ {
    lastAtr := atrs[len(atrs)-1]
    atr := ((lastAtr*float64(period-1)) + tr[i]) / float64(period)
    atrs = append(atrs, atr)
  }
  return atrs
}

func TrueRange(values []HighLowClose) []float64 {
  var result []float64
  for idx,hlc := range values[1:] {
    tr := max([]float64 { (hlc.High-hlc.Low), math.Abs(hlc.High - values[idx].Close), math.Abs(hlc.Low - values[idx].Close) })
    result = append(result, tr)
  }
  return result
}

func max(values []float64) float64 {
  var curMax, prevMax float64
  prevMax = values[0]
  for idx,el := range values[1:] {
    curMax = math.Max(values[idx], el)
    curMax = math.Max(curMax, prevMax)
    prevMax = curMax
  }
  return curMax
}
