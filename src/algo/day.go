package Algo

import (
    "fmt"
    "time"
)

func Day(dateStr string) {
    layout := "02/01/2006" // dd/mm/yyyy format
    date, err := time.Parse(layout, dateStr)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(date.Weekday())
}