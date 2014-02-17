package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt)Error() string {
    message := ""
    if e < 0 {
        message = fmt.Sprintf("負の値:%f", e)
    }
    return message
}
func Sqrt(f float64) (float64, ErrNegativeSqrt) {
    return math.Sqrt(f), ErrNegativeSqrt(f)
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
