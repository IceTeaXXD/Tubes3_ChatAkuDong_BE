package Algo

func border(pattern string) []int {
	b := make([]int, len(pattern))
	b[0] = 0

	var m int = len(pattern)
	var j, i int = 0, 1

	for i < m {
		if pattern[j] == pattern[i] {
			b[i] = j + 1
			i++
			j++
		} else if j > 0 {
			j = b[j-1]
		} else {
			b[i] = 0
			i++
		}
	}

	return b
}

func KMP(pattern string, text string) int {
	var n int = len(text)
	var m int = len(pattern)

	var b []int = border(pattern)

	var i, j int = 0, 0

	for i < n {
		if pattern[j] == text[i] {
			if j == m-1 {
				return i - m + 1 // match
			}
			i++
			j++
		} else if j > 0 {
			j = b[j-1]
		} else {
			i++
		}
	}
	return -1 // no match
}
