package xrand

func randomBig(max int) (random *big.Int, err error) {
	return rand.Int(rand.Reader, big.NewInt(int64(max)))
}