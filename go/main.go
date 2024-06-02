package main

import (
    "fmt"
    "time"
)

func main() {
    startTime := time.Now()

    for i := 1; i <= 1000000; i++ {
        fmt.Println(i)
    }

    duration := time.Since(startTime)
    fmt.Println("Execution time:", duration)
}
