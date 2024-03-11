package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestTimestamps(t *testing.T) {
	fmt.Println(time.Now().UTC())
	// 2022-12-14 20:28:42.72878716 +0000 UTC
}

func TestMultipleDuration(t *testing.T) {
	d := 60
	fmt.Println(time.Duration(d) * time.Second)
}

type Entity struct {
	Name string    `json:"name"`
	Time time.Time `json:"time"`
}

func TestUnmarshalTime(t *testing.T) {
	jsonString := `{"name": "A name", "time": "2021-02-18T21:54:42.123Z"}`
	var entity Entity
	err := json.Unmarshal([]byte(jsonString), &entity)
	if err != nil {
		t.Errorf("failed to unmarshal: %s", err)
	}
	fmt.Println(entity)

}
