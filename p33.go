package main

import "fmt"

func main() {
    var z []int
    fmt.Println(z, len(z), cap(z))
    if z == nil {
        fmt.Println("z is nil!")
    }
}

// TODO: @see http://golang.org/doc/articles/slices_usage_and_internals.html
