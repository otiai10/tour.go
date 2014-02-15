package main

import "fmt"

func fibonacci() func() int {
    former, latter := 0,1
    return func() int {
        sum := former + latter
        former = latter
        latter = sum
        return sum
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
