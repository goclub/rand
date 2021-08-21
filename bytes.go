package xrand
// 基于 seed 返回指定数量的 []byte,主要 seed 中不要包含中文,包含中文请使用 RunesBySeed
func BytesBySeed(seed []byte, size int) ([]byte, error) {
	var result []byte
	for i:=0; i<size; i++ {
		randIndex, err := RangeUint64(0, uint64(len(seed)-1)) ; if err != nil {
			return nil, err
		}
		result = append(result, seed[randIndex])
	}
	return result, nil
}

type Seed struct {}
func (Seed) Alphabet() []byte {
	return []byte("abcdefghijklmnopqrstuvwxyz")
}
func (Seed) Digit() []byte {
	return []byte("0123456789")
}
func (Seed) AlphabetDigit() []byte {
	return []byte("abcdefghijklmnopqrstuvwxyz0123456789")
}