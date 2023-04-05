package xrand

// RunesBySeed
// 基于 seed 返回指定数量的 []rune
func RunesBySeed(seed []rune, size int) []rune {
	var result []rune
	for i := 0; i < size; i++ {
		randIndex := RangeUint64(0, uint64(len(seed)-1))
		result = append(result, seed[randIndex])
	}
	return result
}
