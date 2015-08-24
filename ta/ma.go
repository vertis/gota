package ta

import (
  "fmt"
)

func Mean(values []interface{}) float64 {
  var total float64=0
	for _,element := range values {
		total += element.(float64)
	}
  return total / float64(len(values))
}

func Sma(values []interface{}, period int) []interface{}{
  var result []interface{} = make([]interface{}, len(values))
  for index,_ := range values {
    indexPlusOne := index+1
    if(indexPlusOne>=period) {
      //fmt.Println(index, period, values[indexPlusOne-period:indexPlusOne])
      y := Mean(values[indexPlusOne-period:indexPlusOne])
      result[index] = y
    } else {
      result[index] = nil
    }
  }
  return result
}

func Ema(values []interface{}, period int) []interface{}{
  sma := Sma(values, period)
  var result []interface{} = make([]interface{}, len(values))
  var multiplier = (2.0 / (float64(period) + 1.0) )
  for index,element := range values {
    if(sma[index]==nil || result[index-1]==nil) {
      result[index] = sma[index]
    } else {
      fmt.Println(element.(float64), result[index-1].(float64), (element.(float64) - result[index-1].(float64)))
      result[index] = (element.(float64) - result[index-1].(float64)) * multiplier + result[index-1].(float64)
    }
  }
  fmt.Println(period, multiplier, result)
  return result
}
