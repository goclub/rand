package xrand

import (
	"log"
	"testing"
)

func TestRunesBySeed(t *testing.T) {
	data, err := RunesBySeed([]rune("ab我"), 100)
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	count := map[string]int{}
	if len(data) != 100 {
		log.Print("len error")
		t.Fail()
	}
	for _, item := range data {
		count[string(item)]++
	}
	if count["a"] == 0 {
		log.Print("a")
		t.Fail()
	}
	if count["b"] == 0 {
		log.Print("a")
		t.Fail()
	}
	if count["我"] == 0 {
		log.Print("我")
		t.Fail()
	}
	log.Print(string(data))

}