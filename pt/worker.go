package main

import (
	"log"
	"sync"
	"time"

	"github.com/antonioshadji/goquotemonitor/db"
)

func dowork(w db.Work) {
	bidreq := QuoteRequest{}
	bidreq.Data.Type = "quotes"
	bidreq.Data.Attributes.AccountID = db.Account
	bidreq.Data.Attributes.AssetID = db.Assets[w.Ticker]
	bidreq.Data.Attributes.TransactionType = "sell"
	bidreq.Data.Attributes.UnitCount = w.Size
	if w.LP == "Enigma" {
		bidreq.Data.Attributes.DelayedSettlement = true
		bidreq.Data.Attributes.TradeDeskID = db.TradeDesk["Enigma"]
	}

	askreq := bidreq
	askreq.Data.Attributes.TransactionType = "buy"

	var wg sync.WaitGroup
	ch := make(chan *QuoteResponse, 2)

	// run loop only once per minute
	var delay = 60
	for {
		go func() {
			wg.Add(1)
			defer wg.Done()

			r, err := ptQuoteRequest(&bidreq)
			if err != nil {
				log.Println(err)
			}
			ch <- r
		}()

		go func() {
			wg.Add(1)
			defer wg.Done()

			r, err := ptQuoteRequest(&askreq)
			if err != nil {
				log.Println(err)
			}
			ch <- r
		}()

		wg.Wait()
		// not using a loop so I don't need to close the channel
		// close(ch)

		response1 := <-ch
		response2 := <-ch

		if response1 == nil || response2 == nil {
			// http requests failed with error message
			log.Println("Received nil response from API. No data recorded.")
			time.Sleep(time.Duration(delay) * time.Second)
			continue
		}

		var bid, ask float64
		if response1.Data.Attributes.TransactionType == "buy" {
			ask = response1.Data.Attributes.PricePerUnit
			bid = response2.Data.Attributes.PricePerUnit
		} else {
			ask = response2.Data.Attributes.PricePerUnit
			bid = response1.Data.Attributes.PricePerUnit
		}
		bps := 10000 * ((ask - bid) / bid)

		data := db.InsertRow{
			TS:     response1.Data.Attributes.CreatedAt,
			Bid:    bid,
			Ask:    ask,
			Size:   w.Size,
			Width:  bps,
			Ticker: w.Ticker,
			LP:     w.LP,
		}
		result, err := db.InsertData(data)
		if err != nil {
			log.Println(err)
		}
		ra, _ := result.RowsAffected()
		log.Printf("Inserted %d row for LP: %s Ticker: %s Size: %f.\n", ra, w.LP, w.Ticker, w.Size)
		time.Sleep(time.Duration(delay) * time.Second)
	}
}
