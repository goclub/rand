package xrand

import (
	"crypto/rand"
	"math/big"
)

func Int64(max int64) (random int64, err error) {
	bigInt, err := rand.Int(rand.Reader, big.NewInt(max)) ; if err != nil {
		return 0, err
	}
	return bigInt.Int64(), err
}
