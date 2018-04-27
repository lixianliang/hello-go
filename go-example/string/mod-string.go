package main

import (
    "fmt"
    "strings"
    //"unicode/utf8"
)


func main() {
    s1 := "abcdefg"
    s2 := s1[:2] + "1" + s1[3:]
    fmt.Println(s2)

    s1 = "中国共产党万岁1"
    s2 = strings.Replace(s1, "万", "千", 1)
    fmt.Println(s2)
    //utf8.DecodeRuneInString(s1)
}
