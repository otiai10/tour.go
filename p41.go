package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    res := make(map[string]int)
    for _, word := range strings.Fields(s) {
        res[word] = res[word] + 1
    }
    return res
}

func main() {
    wc.Test(WordCount)
}
