package Algo

import "math"

/**
 * Tabel Last Occurence dari ASCII (128 karakter)
 */
func lastOccurence(text string) [128]int {
	var last [128]int

	for i := 0; i < 128; i++ {
		last[i] = -1
	}

	for i := 0; i < len(text); i++ {
		last[text[i]] = i
	}

	return last
}

func BMMatch(pattern string, text string) int {
	var last [128]int = lastOccurence(text)
	n := len(text)
	m := len(pattern)
	i := m - 1
	if i > n-1 {
		return -1
	} else {
		j := m - 1
		for {
			if pattern[j] == text[i] {
				if j == 0 {
					return i
				} else {
					i--
					j--
				}
			} else {
				var lo int = last[text[i]]
				i = i + m - int(math.Min(float64(j), float64(1+lo)))
				j = m - 1
			}

			if i > n-1 {
				break
			}
		}
	}

	return -1
}
