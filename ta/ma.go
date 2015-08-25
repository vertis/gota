package ta

import (
//  "fmt"
)

func Mean(values []float64) float64 {
  var total float64=0
	for _,element := range values {
		total += element
	}
  return total / float64(len(values))
}

func Sma(values []float64, period int) []float64{
  var result []float64
  for index,_ := range values {
    indexPlusOne := index+1
    if(indexPlusOne>=period) {
      avg := Mean(values[indexPlusOne-period:indexPlusOne])
      result = append(result, avg)
    }
  }
  return result
}

func Wma(values []float64, period int) []float64{
  var result []float64
  for index,_ := range values {
    indexPlusOne := index+1
    if(indexPlusOne>=period) {
      var res []float64
      sl := values[indexPlusOne-period:indexPlusOne]

      // Get the sum of the number of entries
      var sum float64=0
      for idx,_ := range sl {
        sum += float64(idx+1)
      }

      for idx,element := range sl {
        res = append(res, element*(float64(idx+1)/sum))
      }
      var total float64=0
      for _,element := range res {
        total += element
      }
      result = append(result, total)
    }
  }
  return result
}

func Ema(values []float64, period int) []float64 {
  sma := Sma(values, period)
  var result []float64
  var multiplier = (2.0 / (float64(period) + 1.0) )

  result = append(result,sma[0])
  for i := (len(values)-len(sma))+1; i < len(values); i++ {
    lastVal := result[len(result)-1]
    ema := (values[i] - lastVal) * multiplier + lastVal
    result = append(result,ema)
	}
  return result
}

func Dema(values []float64, period int) []float64 {
  var result []float64
  ema := Ema(values, period)
  emaAgain := Ema(ema, period)
  var emaDouble []float64
  for _,element := range ema {
      emaDouble = append(emaDouble, (2.0 * element))
  }
  offset := len(emaDouble)-len(emaAgain)
  for index,element := range emaAgain {
      result = append(result, (emaDouble[index+offset] - element))
  }
  return result
}
