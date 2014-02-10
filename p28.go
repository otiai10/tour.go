package main

import "fmt"

type Vertex struct {
    X, Y int
}

var (
    p = Vertex{1, 3}
    q = &Vertex{1, 3}
    r = Vertex{X: 1}
    s = Vertex{}
)

func main() {
    fmt.Println(p,q,r,s)
}
