package sfox

// TODO: this is duplicate code from coinbase.go

import (
	"fmt"
	"strconv"
)

func CalcSpread(data *BookResponse, size float64) (float64, float64, float64) {

	bid := GetPriceForSize(data.Bids, size)
	ask := GetPriceForSize(data.Asks, size)
	// debug
	//fmt.Println(bid, ask)

	return bid, ask, 10000 * (ask - bid) / bid

}

func GetPriceForSize(data [][]interface{}, size float64) float64 {

	// fmt.Println(len(data), ",psum", ",ssum", ",psum/ssum", ",p", ",s")
	var p, s, psum, ssum, z float64
	var err error
	for i := range data {
		if s, err = strconv.ParseFloat(fmt.Sprintf("%v", data[i][1]), 64); err == nil {
			// fmt.Printf("%T: %f -> ", s, s)
			ssum += s
		}
		if p, err = strconv.ParseFloat(fmt.Sprintf("%v", data[i][0]), 64); err == nil {
			if ssum >= size {
				z = size - (ssum - s)
				ssum = ssum - s + z
				psum += p * z
				break
			}
			psum += p * s
		}
		// debug
		// fmt.Printf("%d, %.8f,%.8f,%.8f,%.8f,%.8f\n", i, psum, ssum, psum/ssum, p, s)
	}
	// fmt.Printf("99, %.8f,%.8f,%.8f,%.8f,%.8f\n", psum, ssum, psum/ssum, p, z)

	return psum / ssum
}
