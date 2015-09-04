package ta

import (
  "math"
)

// Calculates the directional movement. Returns (diffM, diffP)
func directionalMovement(curHigh float64, curLow float64, prevHigh float64, prevLow float64) (float64, float64) {
  diffP    := curHigh-prevHigh /* Plus Delta */
  diffM    := prevLow-curLow   /* Minus Delta */
  if (diffM > 0) && (diffP < diffM) {
      /* Case 2 and 4: +DM=0,-DM=diffM */
      return diffM, 0.0
  } else if (diffP > 0) && (diffP > diffM) {
      return 0.0, diffP
  }
  return 0.0, 0.0 // No change
}

// Adx calculates the Average directional movement index of an array of HighLowClose structs given a period
func Adx(values []HighLowClose, period int) []float64 {
  lookbackTotal := (2 * period) - 1
  startIdx := lookbackTotal
  today := startIdx

  prevMinusDM := 0.0
  prevPlusDM  := 0.0
  prevTR      := 0.0
  today       = startIdx - lookbackTotal
  prevHigh    := values[today].High
  prevLow     := values[today].Low
  prevClose   := values[today].Close
  for i := 0; i < period-1; i++ {
     /* Calculate the prevMinusDM and prevPlusDM */
     today++

     curHigh := values[today].High
     curLow := values[today].Low

     diffM, diffP := directionalMovement(curHigh, curLow, prevHigh, prevLow)
     prevMinusDM += diffM
     prevPlusDM += diffP
     prevHigh = curHigh
     prevLow  = curLow

     truerange := TrueRange(prevHigh,prevLow,prevClose)
     prevTR += truerange
     prevClose = values[today].Close
  }

  /* Add up all the initial DX. */
  sumDX := 0.0
  for i := 0; i < period; i++ {
     /* Calculate the prevMinusDM and prevPlusDM */
     today++

     prevMinusDM -= prevMinusDM/float64(period)
     prevPlusDM  -= prevPlusDM/float64(period)

     curHigh := values[today].High
     curLow := values[today].Low

     diffM, diffP := directionalMovement(curHigh, curLow, prevHigh, prevLow)
     prevMinusDM += diffM
     prevPlusDM += diffP
     prevHigh = curHigh
     prevLow  = curLow

     /* Calculate the prevTR */
     truerange := TrueRange(prevHigh,prevLow,prevClose)
     prevTR = prevTR - (prevTR/float64(period)) + truerange
     prevClose = values[today].Close

     /* Calculate the DX. The value is rounded (see Wilder book). */
     if prevTR != 0.0 {
        minusDI := 100.0*(prevMinusDM/prevTR)
        plusDI  := 100.0*(prevPlusDM/prevTR)
        /* This loop is just to accumulate the initial DX */
        sumDI := minusDI+plusDI
        if sumDI != 0.0 {
           sumDX  += (100.0 * (math.Abs(minusDI-plusDI)/sumDI))
         }
     }
  }

  /* Calculate the first ADX */
  prevADX := sumDX / float64(period)

  /* Output the first ADX */
  var results []float64
  results = append(results, prevADX)

  /* Calculate and output subsequent ADX */
  for  ; today < len(values)-1; {
     /* Calculate the prevMinusDM and prevPlusDM */
     today++

     prevMinusDM -= prevMinusDM/float64(period)
     prevPlusDM  -= prevPlusDM/float64(period)

     curHigh := values[today].High
     curLow := values[today].Low

     diffM, diffP := directionalMovement(curHigh, curLow, prevHigh, prevLow)
     prevMinusDM += diffM
     prevPlusDM += diffP
     prevHigh = curHigh
     prevLow  = curLow

     /* Calculate the prevTR */
     truerange := TrueRange(prevHigh,prevLow,prevClose)
     prevTR = prevTR - (prevTR/float64(period)) + truerange
     prevClose = values[today].Close

     if prevTR != 0.0 {
        /* Calculate the DX. The value is rounded (see Wilder book). */
        minusDI  := 100.0*(prevMinusDM/prevTR)
        plusDI   := 100.0*(prevPlusDM/prevTR)
        sumDI := minusDI+plusDI
        if sumDI!=0.0 {
           dx := 100.0*(math.Abs(minusDI-plusDI)/sumDI)
           /* Calculate the ADX */
           prevADX = ((prevADX*(float64(period)-1))+dx)/float64(period)
        }
     }

     /* Output the ADX */
     results = append(results, prevADX)
  }

  return results
}
