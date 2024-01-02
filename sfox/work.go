package sfox

import (
	"log"

	"github.com/antonioshadji/goquotemonitor/db"
)

func Work(w db.Work) {
	log.Println(w)
}
