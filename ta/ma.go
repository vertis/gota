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
      fmt.Println(index, period, values[indexPlusOne-period:indexPlusOne])
      y := Mean(values[indexPlusOne-period:indexPlusOne])
      result[index] = y
    } else {
      result[index] = nil
    }
  }
  return result
}
