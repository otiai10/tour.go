package main

import (
    "fmt"
)

func main() {
    c := complex(1, 2)
    fmt.Println(
        c,
    )
    c_2 := c**2
    fmt.Println(c_2)
}
