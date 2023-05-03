package Algo

func MatchRatio(s1, s2 string) float64 {
    if len(s1) > len(s2) {
        s1, s2 = s2, s1
    }
    len1, len2 := len(s1), len(s2)
    if len2 == 0 {
        return 1.0
    }
    row1 := make([]int, len2+1)
    for i := 0; i <= len2; i++ {
        row1[i] = i
    }
    for i := 0; i < len1; i++ {
        row2 := make([]int, len2+1)
        row2[0] = i + 1
        for j := 0; j < len2; j++ {
            cost := 1
            if s1[i] == s2[j] {
                cost = 0
            }
            row2[j+1] = min(row2[j]+1, row1[j+1]+1, row1[j]+cost)
        }
        row1 = row2
    }
    return 1.0 - float64(row1[len2])/float64(max(len1, len2))
}

func min(a, b, c int) int {
    if a < b {
        if a < c {
            return a
        }
        return c
    }
    if b < c {
        return b
    }
    return c
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}