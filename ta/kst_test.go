package ta_test

import (
	. "github.com/vertis/gota/ta"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	//"encoding/csv"
	"io/ioutil"
	"strings"
	"strconv"
	//"fmt"
	"math"
)

type YahooFinanceData struct {
	Date string
	Open float64
	High float64
	Low float64
	Close float64
	Volume int64
}

func getFixtureData(path string) []YahooFinanceData {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data),"\n")
	//_ = strings.Split(lines[0], ",")
	var records []YahooFinanceData
	for _,el := range lines {
		cols := strings.Split(el, ",")

		if len(cols) > 1 {
			open, _ := strconv.ParseFloat(cols[1], 64)
			high, _ := strconv.ParseFloat(cols[2], 64)
			low, _ := strconv.ParseFloat(cols[3], 64)
			close, _ := strconv.ParseFloat(cols[4], 64)

			vol, _ := strconv.ParseInt(cols[5], 10, 64)
			tmp := YahooFinanceData{Date: cols[0], Open: open, High: high, Low: low, Close: close, Volume: vol}
			records = append(records, tmp)
		}
	}

	return records
}

func getResultData() ([]float64, []float64){
	data, err := ioutil.ReadFile("../fixtures/kst-results-ruby.csv")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data),"\n")
	//_ = strings.Split(lines[0], ",")
	var ksts []float64
	var signals []float64
	for _,el := range lines {
		cols := strings.Split(el, ",")
		if len(cols) > 1 {
			kst, err := strconv.ParseFloat(cols[0], 64)
			if err == nil {
				ksts = append(ksts, kst)
			}
			signal, err := strconv.ParseFloat(cols[1], 64)
			if err == nil {
				signals = append(signals, signal)
			}
		}
	}
	return ksts, signals
}

// https://gist.github.com/DavidVaini/10308388
func round(val float64, roundOn float64, places int ) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func roundedValues(values []float64, roundOn float64, places int) ([]float64) {
	var roundedVals []float64
	for _,el := range values {
		roundedVals = append(roundedVals, round(el, roundOn, places))
	}
	return roundedVals
}

var _ = Describe("Kst", func() {
	It("should return the correct values", func() {
		data := getFixtureData("../fixtures/AAPL.csv")
		var closes []float64
		for _,el := range data {
			closes = append(closes, el.Close)
		}

		kst_line_out, kst_signal_line_out := getResultData()
		kst_line, kst_signal_line := Kst(closes, []int { 10, 15, 20, 30}, []int {10, 10, 10, 15}, 9)

		Expect(roundedValues(kst_line, .5, 6)).To(Equal(roundedValues(kst_line_out, .5, 6)))
		Expect(roundedValues(kst_signal_line, .5, 6)).To(Equal(roundedValues(kst_signal_line_out, .5, 6)))
	})
})
