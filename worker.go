package main

import (
	"log"
	"sync"
	"time"
)

func dowork(w Work) {
	bidreq := PTQuotesRequestBody{}
	bidreq.Data.Type = "quotes"
	bidreq.Data.Attributes = QuoteRequestAttrs{
		AccountID:       account,
		AssetID:         assets[w.ticker],
		TransactionType: "sell",
		UnitCount:       w.size,
	}
	if w.lp == "Enigma" {
		bidreq.Data.Attributes.DelayedSettlement = true
		bidreq.Data.Attributes.TradeDeskID = tradeDesk["Enigma"]
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
		data := InsertRow{
			response1.Data.Attributes.CreatedAt,
			bid,
			ask,
			w.size,
			bps,
			w.ticker,
			w.lp,
		}
		result, err := insertData(data)
		if err != nil {
			log.Println(err)
		}
		ra, _ := result.RowsAffected()
		log.Printf("Inserted %d row for LP: %s ticker: %s size: %f.\n", ra, w.lp, w.ticker, w.size)
		time.Sleep(time.Duration(delay) * time.Second)
	}
}
