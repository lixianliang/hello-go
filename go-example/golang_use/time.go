package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.Now()

    fmt.Printf("%#v\n", t)
    fmt.Printf("ts: %d\n", t.Unix())
    str, off := t.Zone()
    fmt.Printf("zone: %s, offset:%d\n", str, off)

    b, _ := t.MarshalBinary()
    fmt.Printf("%b\n", b)

    bjson, _ := t.MarshalJSON()
    fmt.Printf("%s\n", bjson)

    btext, _ := t.MarshalText()
    fmt.Printf("%s\n", btext)

    t1 := time.Now()
    time.Sleep(3*time.Second)
    t2 := time.Now()
    fmt.Printf("after: %t\n", t1.After(t2))
    fmt.Printf("before: %t\n", t1.Before(t2))
    fmt.Printf("equal: %t\n", t1.Equal(t2))
    fmt.Printf("month: %s\n", time.January.String())
    fmt.Printf("day: %s\n", time.Sunday.String())

    y, m, d := t1.Date()
    fmt.Printf("%d %d %d\n", y, m, d)
    // t.Year t.Month t.Day t.Weekday
    h, f, s := t1.Clock()
    fmt.Printf("%d %d %d\n", h, f, s)
    // t.Hour t.Minute t.Second t.Nanosecond
    var dur time.Duration;
    dur = 4000
    fmt.Printf("dur: %s\n", dur.String())
    // t.Add(d)
    // t.Sub(u) 时间相差
    // Since
    // AddDate
    // 
    // Unix(sec, nsec) Time
    // Date() Time
    fmt.Printf("time layout default: %s\n", time.Now().String())
    fmt.Printf("time format layout %s\n", time.Now().Format("2006:01:02 15:04:05"))
    fmt.Printf("time ansic %s\n", time.Now().Format(time.ANSIC))

    xt := time.Now()
    pt, _ := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", xt.String())
    // utc
    plt, err1 := time.ParseInLocation("2006-01-02 15:04:05.999999999 -0700 MST", xt.String(), xt.Location())
    if err1 != nil {
        fmt.Printf("err1: %v", err1)
    }
    plt2, err2 := time.ParseInLocation("2006-01-02 15:04:05.999999999 -0700 MST", xt.String(), time.UTC)
    if err2 != nil {
        fmt.Printf("err2: %v", err2)
    }
    fmt.Printf("1: %d 2: %d 3: %d 4: %d", xt.Unix(), pt.Unix(), plt.Unix(), plt2.Unix())
}
