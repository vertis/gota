package ta

import (
  //"math"
  //"fmt"
)

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

// func Kst(values []float64, rocPeriods []int, avgPeriods []int, signalPeriod int) ([]float64,[]float64) {
//   // TODO: Validate that I'm getting 4 roc periods and 4 avgPeriods
//   maxRocPeriod := 0
//   for _,el := range rocPeriods {
//     maxRocPeriod = int(math.Max(float64(maxRocPeriod), float64(el)))
//   }
//   maxAvgPeriod := 0
//   for _,el := range avgPeriods {
//     maxAvgPeriod = int(math.Max(float64(maxAvgPeriod), float64(el)))
//   }
//   var ksts []float64
//   var kstsignals []float64
//   for idx,_ := range values[maxRocPeriod+maxAvgPeriod-1:] {
//     valuesIdx := idx+maxRocPeriod+maxAvgPeriod+1
//     fmt.Println(valuesIdx)
//     roc1 := Roc(values[:valuesIdx],rocPeriods[0])
//     roc2 := Roc(values[:valuesIdx],rocPeriods[1])
//     roc3 := Roc(values[:valuesIdx],rocPeriods[2])
//     roc4 := Roc(values[:valuesIdx],rocPeriods[3])
//     //fmt.Println(roc1, roc2, roc3, roc4)
//     sma1 := Sma(roc1,avgPeriods[0]) //Mean(roc1[len(roc1)-avgPeriods[0]:])
//     sma2 := Sma(roc2,avgPeriods[1]) //Mean(roc2[len(roc2)-avgPeriods[1]:])
//     sma3 := Sma(roc3,avgPeriods[2]) //Mean(roc3[len(roc3)-avgPeriods[2]:])
//     sma4 := Sma(roc4,avgPeriods[3]) //Mean(roc4[len(roc4)-avgPeriods[3]:])
//     //fmt.Println(sma1, sma2, sma3, sma4)
//     kst := (sma1[len(sma1)-1] * 1) + (sma2[len(sma2)-1] * 2) + (sma3[len(sma3)-1] * 3) + (sma4[len(sma4)-1] * 4)
//     ksts = append(ksts, kst)
//     //fmt.Println(kst, len(ksts))
//     if len(ksts) >= signalPeriod {
//       signal := Mean(ksts[len(ksts)-signalPeriod:])
//       kstsignals = append(kstsignals, signal)
//     }
//   }
//   return ksts, kstsignals
// }
