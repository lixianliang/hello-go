package main 

import (
    "bufio"
    "bytes"
    "fmt"
)

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
    scanner := bufio.NewScanner(bytes.NewReader(p))
    scanner.Split(bufio.ScanLines)
    count := 0
    for scanner.Scan() {
        count++
    }
    if err := scanner.Err(); err != nil {
        fmt.Println("reading input: ", err)
        return 0, err
    }
    //fmt.Println(count)
    *c = LineCounter(count)
    return count, nil
}

func main() {
    var c LineCounter
    c.Write([]byte("a bb ccc \ndddd"))
    fmt.Println(c)

    c = 0
    var name = "li\n xian\n liang\n"
    fmt.Fprintf(&c, "hello\n %s kk", name)
    fmt.Println(c)
    fmt.Printf("%T\n", c)
}
