package xrand
// 基于 seed 返回指定数量的 []rune
func RunesBySeed(seed []rune, size int) ([]rune, error) {
	var result []rune
	for i:=0; i<size; i++ {
		randIndex, err := Int64(int64(len(seed))) ; if err != nil {
      return nil, err
    }
		result = append(result, seed[randIndex])
	}
	return result, nil
}
