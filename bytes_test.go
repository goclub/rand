package xrand

import (
	"log"
	"testing"
)

func TestBytesBySeed(t *testing.T) {
	data := BytesBySeed([]byte("abc"), 100)
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
		log.Print("b")
		t.Fail()
	}
	if count["c"] == 0 {
		log.Print("c")
		t.Fail()
	}
}
