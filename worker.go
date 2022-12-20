package main

import (
	"log"
	"sync"
	"time"
)

var tradeDesk = map[string]string{
	"DV":     "2fafbc29-d5a1-4a68-a8c6-cf93b043ad1b",
	"Enigma": "4c248890-703d-4ac3-9ce7-9de6465f328a",
}

var assets = map[string]string{
	"AVAX": "883a179b-91df-4b42-a31f-02b6c9736537",
	"BTC":  "43454116-7026-4b3d-a9de-eeaada500d4c",
	"ETH":  "efd4e846-99ec-4362-b21f-7982529bf570",
	"LTC":  "b2a5a8e2-d9d5-4854-8517-5a5608ffbe7d",
	"SOL":  "086d6588-6dec-4807-97c2-82e97ef6ff10",
	"USDC": "70c21c75-9362-4182-a984-daa63e30ee52",
	"USDT": "25311c35-26d0-4cf1-b916-80f377a7e468",
}

var account = "0c7715e3-7cdd-4d49-88bb-f1ab3cb8803b"

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
		bidreq.Data.Attributes.TradeDeskID = "4c248890-703d-4ac3-9ce7-9de6465f328a"
	}

	askreq := bidreq
	askreq.Data.Attributes.TransactionType = "buy"

	var wg sync.WaitGroup
	ch := make(chan *QuoteResponse, 2)

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
			time.Sleep(3 * time.Second)
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
		time.Sleep(60 * time.Second)
	}
}
