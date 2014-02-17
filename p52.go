package main

import (
    "fmt"
    "math"
)

type Verrex struct {
    X, Y float64
}

func (v *Verrex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func (v *Verrex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
    v := &Verrex{3,4}
    v.Scale(5)
    fmt.Println(v, v.Abs())
}
