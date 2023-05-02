package Algo

import (
    "strconv"
)

func Day(dateStr string) int{
    /* Input dateStr dalam bentuk dd/mm/yyyy */
    day := dateStr[:2]
    d, _ := strconv.Atoi(day)
    month := dateStr[3:5]
    m, _ := strconv.Atoi(month)
    year := dateStr[6:]
    y, _ := strconv.Atoi(year)
    
    return getDay(y, m, d)
}

func getDay(year, month, day int) int {
    t := []int{0, 3, 2, 5, 0, 3, 5, 1, 4, 6, 2, 4}
    if month < 3 {
        year -= 1
    }
    return (year + year/4 - year/100 + year/400 + t[month-1] + day) % 7
}