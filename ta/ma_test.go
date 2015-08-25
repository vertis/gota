package ta

import "testing"

func TestSma(t *testing.T) {
	cases := []struct {
		in []float64
    want float64
	}{
		{ []float64{ 0.0,1.0,2.0,3.0,4.0,5.0,6.0,7.0,8.0,9.0 }, 7.0 },
    { []float64{ 0.0,1.0,2.0,3.0,4.0,5.0,6.0,7.0,8.0 }, 6.0 },
	}
	for _, c := range cases {
    var interfaceSlice []interface{} = make([]interface{}, len(c.in))
    for i, d := range c.in {
      interfaceSlice[i] = d
    }
		got := Sma(interfaceSlice, 5)
		if got[len(got)-1] != c.want {
			t.Errorf("Sma(%q) == %q, want %q", c.in,got[len(got)-1], c.want)
		}
	}
}

func TestEma(t *testing.T) {
	cases := []struct {
		in []float64
    want float64
	}{
		{ []float64{ 26.0, 54.0, 8.0, 77.0, 61.0, 39.0, 44.0, 91.0, 98.0, 17.0 }, 53.7917695473251 },
	}
	for _, c := range cases {
    var interfaceSlice []interface{} = make([]interface{}, len(c.in))
    for i, d := range c.in {
      interfaceSlice[i] = d
    }
		got := Ema(interfaceSlice, 5)
		if got[len(got)-1] != c.want {
			t.Errorf("Ema(%q) == %q, want %q", c.in,got[len(got)-1], c.want)
		}
	}
}

func TestDema(t *testing.T) {
	cases := []struct {
		in []float64
    want float64
	}{
		{ []float64{ 26.0, 54.0, 8.0, 77.0, 61.0, 39.0, 44.0, 91.0, 98.0, 17.0 }, 54.55632373113854 }, //[91.73037037037035, 54.55632373113854, nil, nil, nil, nil, nil, nil, nil, nil]
	}
	for _, c := range cases {
    var interfaceSlice []interface{} = make([]interface{}, len(c.in))
    for i, d := range c.in {
      interfaceSlice[i] = d
    }
		got := Dema(interfaceSlice, 5)
		if got[len(got)-1] != c.want {
			t.Errorf("Dema(%q) == %q, want %q", c.in,got[len(got)-1], c.want)
		}
	}
}
