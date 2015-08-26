package ta

import (
  "math"
)

// Rsi calculates the Relative Strength Index for an
// array of []float64 numbers and a given period
func Rsi(values []float64, period int) []float64 {
    var up []float64
    var down []float64
    for index,element := range values[1:] {
      previous := values[index]
      if(element >= previous) {
        up = append(up, element-previous)
        down = append(down, 0)
      } else {
        down = append(down, previous-element)
        up = append(up, 0)
      }
    }
    var rs []float64
    var rsi []float64
    prevAvgUp := Mean(up[0:period-1])
    prevAvgDown := Mean(down[0:period-1])
    for i := period-1; i < len(up); i++{
      currAvgUp := (prevAvgUp * float64(period-1) + up[i]) / float64(period)
      currAvgDown := (prevAvgDown * float64(period-1) + down[i]) / float64(period)
      currRS := currAvgUp/math.Abs(currAvgDown)
      rs = append(rs, currRS)
      currRSI := 100 - 100 / (1 + currRS)
      rsi = append(rsi, currRSI)
      prevAvgUp = currAvgUp
      prevAvgDown = currAvgDown
    }
    return rsi
}
// func Rsi(values []float64, period int) []float64{
//
//   var up []float64
//   var down []float64
//   for index,element := range values[1:] {
//     previous := values[index]
//     if(element >= previous) {
//       up = append(up, element-previous)
//       down = append(down, 0)
//     } else {
//       down = append(down, previous-element)
//       up = append(up, 0)
//     }
//   }
//   // emaUp := Ema(up, period)
//   // emaDown := Ema(down, period)
//   // var rs []float64
//   // for idx,el := range emaUp {
//   //   rs = append(rs, el / emaDown[idx])
//   // }
//   rollingUp := Sma(up,period)
//   rollingDown := Sma(down, period)
//   var rs []float64
//   for idx,el := range rollingUp {
//     rs = append(rs, el / math.Abs(rollingDown[idx]))
//   }
//   var result []float64
//   for _,e := range rs {
//     rsi := 100 - 100 / (1+e)
//     result = append(result, rsi)
//   }
//   return result
// }
