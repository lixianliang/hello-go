package main

import (
    "math/rand"
    "testing"
)

func BenchmarkRandom(b *testing.B) {
    for i := 0; i < b.N; i++ {
        random()
    }
}

func random() int {
    return rand.Intn(100)
}

func main() {
}
