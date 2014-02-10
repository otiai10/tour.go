package main

import (
    "fmt"
    "math"
)

func Sqrt(x, start float64) float64 {
    for i := 0; i<10; i++ {
        start = start - (math.Pow(start,2) - x)/(2 * start)
    }
    return start
}

func main() {
    fmt.Println(
        Sqrt(2, 1),
        Sqrt(2, 2),
        Sqrt(2, 3),
        Sqrt(2, 100),
        Sqrt(2, 1.5),
    )
}
