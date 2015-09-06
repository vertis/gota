package ta

import (
  //"math"
  //"fmt"
)

// Calculates Martin Pring's Know Sure Thing for a slice of float64
func Kst(values []float64, rocPeriods []int, avgPeriods []int, signalPeriod int) ([]float64,[]float64) {
  var ksts []float64
  for i := 0; i < len(values); i++ {
    rocma1 := Sma(Roc(values[:i+1], rocPeriods[0]), avgPeriods[0])
    rocma2 := Sma(Roc(values[:i+1], rocPeriods[1]), avgPeriods[1])
    rocma3 := Sma(Roc(values[:i+1], rocPeriods[2]), avgPeriods[2])
    rocma4 := Sma(Roc(values[:i+1], rocPeriods[3]), avgPeriods[3])
    if !(len(rocma1)>0 && len(rocma2)>0 && len(rocma3)>0 && len(rocma4)>0) {
      continue
    }
    kst := rocma1[len(rocma1)-1] + (rocma2[len(rocma2)-1]*2) + (rocma3[len(rocma3)-1]*3) + (rocma4[len(rocma4)-1]*4)
    ksts = append(ksts, kst)
  }
  kstSignals := Sma(ksts, signalPeriod)
  return ksts, kstSignals
}
