package main 

import (
    "bufio"
    "bytes"
    "fmt"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
    scanner := bufio.NewScanner(bytes.NewReader(p))
    scanner.Split(bufio.ScanWords)
    count := 0
    for scanner.Scan() {
        count++
    }
    if err := scanner.Err(); err != nil {
        fmt.Println("reading input: ", err)
        return 0, err
    }
    //fmt.Println(count)
    *c = WordCounter(count)
    return count, nil
}

func main() {
    var c WordCounter
    c.Write([]byte("a bb ccc dddd"))
    fmt.Println(c)

    c = 0
    var name = "li xian liang"
    fmt.Fprintf(&c, "hello %s kk", name)
    fmt.Println(c)
    fmt.Printf("%T\n", c)
}
