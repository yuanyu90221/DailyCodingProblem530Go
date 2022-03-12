package diff_string

func GetStringDistance(str1, str2 string) int {
	str1Len := len(str1)
	str2Len := len(str2)
	diff := 0
	min := str1Len
	max := str2Len
	target := str1
	compareStr := str2
	if str2Len < str1Len {
		min = str2Len
		max = str1Len
		target = str2
		compareStr = str1
	}

	for idx := range target {
		if target[idx] != compareStr[idx] {
			diff++
		}
	}
	return diff + (max - min)
}
