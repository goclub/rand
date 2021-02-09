package xrand
// 基于 seed 返回指定数量的 []rune
func RunesBySeed(seed []rune, size int) ([]rune, error) {
	var result []rune
	for i:=0; i<size; i++ {
		randIndex, err := randomBig(len(seed)) ; if err != nil {
      return nil, err
    }
		result = append(result, seed[randIndex.Int64()])
	}
	return result, nil
}
