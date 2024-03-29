package sfox

// TODO: mostly duplicate with coinbase cbworker.

import (
	"log"
	"time"

	"github.com/antonioshadji/goquotemonitor/db"
)

var vendor = "sFOX"

func Work(w db.Work) {
	for {
		response, err := GetBook(w.Ticker)
		if err != nil {
			log.Printf("Error: %v", err)
			time.Sleep(3 * time.Second)
			continue
		}

		bid, ask, bps := CalcSpread(response, w.Size)

		data := db.InsertRow{
			TS:     time.Now().UTC(),
			Bid:    bid,
			Ask:    ask,
			Size:   w.Size,
			Width:  bps,
			Ticker: w.Ticker,
			LP:     vendor,
		}
		result, err := db.InsertData(data)
		if err != nil {
			log.Println(err)
		}
		ra, _ := result.RowsAffected()
		log.Printf("Inserted %d row for LP: %s ticker: %s size: %f.\n", ra, vendor, w.Ticker, w.Size)

		time.Sleep(60 * time.Second)
	}
}
