package ta

type CloseVolume struct {
  Close float64
  Volume float64
}

// Obv calculates the On Balance Volume indicator
func Obv(values []CloseVolume) []float64 {
  var result []float64
  acc := 0.0
  for idx,el := range values[1:] {
    if values[idx].Close > el.Close {
      acc = acc - el.Volume
      result = append(result, acc)
    } else if values[idx].Close < el.Close {
      acc = acc + el.Volume
      result = append(result, acc)
    } else {
      result = append(result, acc)
    }
  }
  return result
}
