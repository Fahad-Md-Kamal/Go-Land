package twosum

func IsValid(arr *[]int, target int) []int{
	m := make(map[int]int)

	for key, value := range *arr{
		if val, ok := m[target - value]; ok {
			if val != key{
				arr := []int{ val, key}
				return arr
			}
		}

		m[value] = key
	}
	return []int{-1,-1}
}