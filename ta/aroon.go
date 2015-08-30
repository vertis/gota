package ta

import (
  "math"
)

// Aroon calculates the Aroon Indicator for the
// supplied array of float64 values for a given period
// The formula is:
// AroonUp = ((<period> - Days Since <period> High)/<period>) x 100
// AroonDown = ((<period> - Days Since <period> Low)/<period>) x 100
func Aroon(values []float64, period int) ([]float64, []float64) {
  var aroonDowns, aroonUps []float64
  for i := period; i < len(values); i++ {
    currentSegment := values[i-period:i+1]
    maxIdx := maxIndex(currentSegment)
    minIdx := minIndex(currentSegment)
    daysSinceHigh := len(currentSegment)-maxIdx-1
    daysSinceLow := len(currentSegment)-minIdx-1
    aroonDown := ((float64(period) - float64(daysSinceLow))/float64(period)) * 100.0
    aroonUp := ((float64(period) - float64(daysSinceHigh))/float64(period)) * 100.0
    aroonDowns = append(aroonDowns, aroonDown)
    aroonUps = append(aroonUps, aroonUp)
  }
  return aroonDowns, aroonUps
}

func maxIndex(values []float64) int {
  var curMax, prevMax float64
  prevMax = values[0]
  maxIdx := 0
  for idx,el := range values[1:] {
    curMax = math.Max(values[idx], el)
    curMax = math.Max(curMax, prevMax)
    if(curMax > prevMax) {
      maxIdx = idx+1
    }
    prevMax = curMax
  }
  return maxIdx
}

func minIndex(values []float64) int {
  var curMin, prevMin float64
  prevMin = values[0]
  minIdx := 0
  for idx,el := range values[1:] {
    curMin = math.Min(values[idx], el)
    curMin = math.Min(curMin, prevMin)
    if(curMin < prevMin) {
      minIdx = idx+1
    }
    prevMin = curMin
  }
  return minIdx
}
