package kraken

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/antonioshadji/goquotemonitor/db"
)

func TestKrakenQuote(t *testing.T) {
	res, err := getBook("XBTUSD")
	if err != nil {
		t.Errorf("FAIL: %#v", err)
	}
	for k := range res.Result {
		fmt.Printf("Price Quantity: %s %s\n", res.Result[k].Bids[0][0], res.Result[k].Bids[0][1])
		fmt.Printf("Price Quantity: %s %s\n", res.Result[k].Asks[0][0], res.Result[k].Asks[0][1])
		fmt.Printf("Types: %T %T\n", res.Result[k].Asks[0][0], res.Result[k].Asks[0][1])

	if s, err := strconv.ParseFloat(fmt.Sprintf("%v", res.Result[k].Asks[0][0]), 32); err == nil {
		fmt.Printf("%T: %f\n", s, s)
	}
	if s, err := strconv.ParseFloat(fmt.Sprintf("%v", res.Result[k].Asks[0][0]), 64); err == nil {
		fmt.Printf("%T: %f\n", s, s)
	}
	if s, err := strconv.ParseFloat(fmt.Sprintf("%v", res.Result[k].Asks[0][1]), 32); err == nil {
		fmt.Printf("%T: %f\n", s, s)
	}
	if s, err := strconv.ParseFloat(fmt.Sprintf("%v", res.Result[k].Asks[0][1]), 64); err == nil {
		fmt.Printf("%T: %f\n", s, s)
	}
	}
}

func TestCalcSpread(t *testing.T) {
	f, e := os.ReadFile("../testdata/kraken.json")
	if e != nil {
		t.Errorf("file read failed: %#v\n", e)
	}

	var data OrderBookResponse

	e = json.Unmarshal([]byte(f), &data)
	if e != nil {
		t.Errorf("Unmarshal failed: %#v\n", e)
	}

	// fmt.Printf("bids: %v\n", data.Bids)
	// fmt.Printf("asks: %v\n", data.Asks)

	bid, ask, result := calcSpread(&data, 1)
	fmt.Println(bid, ask, result, "bps")
	fmt.Printf("%.3f bps\n", result)

}

func TestWork(t *testing.T) {
	// TODO: include in testing
	t.SkipNow()
	work := db.Work{
		LP:     "Kraken",
		Ticker: "BTC",
		Size:   1.0,
	}

	// this code does not return
	Work(work)
}
