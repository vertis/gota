package main
import (
  "fmt"
  "github.com/vertis/gota/ta"
)

func main() {
  x := []float64{0.0,1.0,2.0,3.0,4.0,5.0,6.0,7.0,8.0,9.0}
  var interfaceSlice []interface{} = make([]interface{}, len(x))
  for i, d := range x {
    interfaceSlice[i] = d
  }
  y := ta.Sma(interfaceSlice, 5)
  fmt.Println(y)
}
