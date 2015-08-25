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

func Ema(values []interface{}, period int) []interface{} {
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

func Dema(values []interface{}, period int) []interface{} {
  var result []interface{} = make([]interface{}, len(values))
  ema := Ema(values, period)
  emaAgain := Ema(ema, period)
  var emaDouble []interface{} = make([]interface{}, len(values))
  for index,element := range ema {
    if(element==nil) {
      emaDouble[index] = nil
    } else {
      emaDouble[index] =  (2.0 * element.(float64))
    }
  }
  for index,element := range emaAgain {
    if(element==nil || emaDouble[index]==nil) {
      result[index] = nil
    } else {
      result[index] =  emaDouble[index].(float64) - element.(float64)
    }
  }
  return result
}
