package xrand
// 基于 seed 返回指定数量的 []rune
func RunesBySeed(seed []rune, size int) ([]rune, error) {
	var result []rune
	for i:=0; i<size; i++ {
		randIndex, err := RangeUint64(0, uint64(len(seed)-1)) ; if err != nil {
		    return nil, err
		}
		result = append(result, seed[randIndex])
	}
	return result, nil
}
