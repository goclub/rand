package xrand

import (
	"crypto/rand"
	"errors"
	"math/big"
)

// RandUint64(0, 2) return 0 or 1 or 2 , RandUint64(2, 3) return 2 or 3
func RangeUint64(min uint64, max uint64) (random uint64, err error) {
	if max == 0 {
		return 0, errors.New("goclub/rand: Int64(max) max can not be 0")
	}
	bigInt, err := rand.Int(rand.Reader, new(big.Int).SetUint64(max+1-min)) ; if err != nil {
		return 0, err
	}
	random = bigInt.Uint64() + min
	return random, err
}

// RandInt64(0, 2) return 0 or 1 or 2 , RandUint64(2, 3) return 2 or 3
func RangeInt64(min int64, max int64) (random int64, err error) {
	if max == 0 {
		return 0, errors.New("goclub/rand: Int64(max) max can not be 0")
	}
	bigInt, err := rand.Int(rand.Reader, new(big.Int).SetInt64(max+1-min)) ; if err != nil {
		return 0, err
	}
	random = bigInt.Int64() + min
	return random, err
}
