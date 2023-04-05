package xrand

import (
	"github.com/stretchr/testify/assert"
	"log"
	"math"
	"reflect"
	"testing"
)

func TestRangeUint64(t *testing.T) {
	random := RangeUint64(0, 0)
	assert.Equal(t, random, uint64(0))
	// 测试最大值不会panic
	log.Print(RangeUint64(0, math.MaxUint64))
	var testRange = func(min uint64, max uint64) {
		result := map[uint64]uint64{}

		for i := 0; i < 100; i++ {
			random := RangeUint64(min, max)
			result[random]++
		}
		keys := reflect.ValueOf(result).MapKeys()
		assert.Equal(t, len(keys), int(max+1-min))
		for i := min; i < max+1; i++ {
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

func TestRangeInt64(t *testing.T) {
	random := RangeInt64(0, 0)
	// 测试最小最大值不会panic
	log.Print(RangeInt64(math.MinInt64, math.MaxInt64))
	assert.Equal(t, random, int64(0))
	var testRange = func(min int64, max int64) {
		result := map[int64]int64{}

		for i := 0; i < 100; i++ {
			random := RangeInt64(min, max)
			result[random]++
		}
		keys := reflect.ValueOf(result).MapKeys()
		assert.Equal(t, len(keys), int(max+1-min))
		for i := min; i < max+1; i++ {
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

	testRange(-2, 1)
	testRange(-2, 2)
	testRange(-2, 3)
	testRange(-2, 4)
	testRange(-2, 5)
	testRange(-2, 6)
	testRange(-2, 1)
	testRange(-2, 2)
	testRange(-2, 3)
	testRange(-2, 2)
	testRange(-2, 3)
	testRange(-2, 4)
}
