package xrand

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestRangeUInt64(t *testing.T) {
	for i:=0;i<100;i++ {
		random, err := RangeUint64(0, 0)
		assert.EqualError(t, err, "goclub/rand: Int64(max) max can not be 0")
		assert.Equal(t, random, uint64(0))
	}
	var testRange = func(min uint64, max uint64) {
		result := map[uint64]uint64{}

		for i:=0;i<100;i++ {
			random, err := RangeUint64(min, max) ; assert.NoError(t, err)
			result[random]++
		}
		keys := reflect.ValueOf(result).MapKeys()
		assert.Equal(t, len(keys), int(max+1-min))
		for i := min;i<max+1;i++ {
			_, has := result[i]
			assert.Equal(t, has, true, "min", min, "max", max)
		}
	}
	testRange(0, 1)
	testRange(0, 2)
	testRange(0, 3)
	testRange(0, 4)
	testRange(0, 5)
	testRange(0, 6)
	testRange(1, 1)
	testRange(1, 2)
	testRange(1, 3)
	testRange(2, 2)
	testRange(2, 3)
	testRange(2, 4)
}
