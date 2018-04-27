package main

import (
	"fmt"
	"log"
    "encoding/json"
)

var data = `
    "abc"
`

func main() {
    var t interface{}
    err := json.Unmarshal([]byte(data), &t)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
	fmt.Printf("--- t:\n%#v\n\n", t)
    xxx(t)

    data = `123`
    err = json.Unmarshal([]byte(data), &t)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
	fmt.Printf("--- t:\n%+v\n\n", t)
    xxx(t)

    data = `[1, 2, 3]`
    err = json.Unmarshal([]byte(data), &t)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
	fmt.Printf("--- t:\n%T\n\n", t)
    xxx(t)

    data = `{"a":1, "b":2, "c":3}`
    err = json.Unmarshal([]byte(data), &t)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
	fmt.Printf("--- t:\n%#v\n\n", t)
    xxx(t)

    data = `{"data":"xxx", "http":{"header":{"cookie":"abc", "requet-id":"123"}, "code":200}}`
    err = json.Unmarshal([]byte(data), &t)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
	fmt.Printf("--- t:\n%#v\n\n", t)
    xxx(t)
    x, ok := t.(map[string]interface {})
    if ok {
        obj, ok := x["http"]
        if ok {
            x, ok = obj.(map[string]interface {})
            if ok {
                obj, ok = x["header"]
                if ok {
                    x = obj.(map[string]interface {})
                    if ok {
                        for k, v := range x {
                            v, ok := v.(string)
                            if ok {
                                fmt.Printf("k:%s v:%s\n", k, v)
                            }
                        } 
                    }
                }
            }
        }
    }
}

func xxx(t interface{}) {
    fmt.Printf("xxx === ");
    switch x := t.(type) {
    case bool:
        fmt.Printf("bool %t\n", t)
    case int:
        fmt.Printf("interger %d\n", x)
    case float64:
        fmt.Printf("float64 %f\n", x)
    case string:
        fmt.Printf("string %s\n", x)
    case []interface {}:
        fmt.Printf("array %v\n", x)
    case map[string]interface {}:
        fmt.Printf("map[string] interface{} %v\n", x)
    default:
        fmt.Printf("unexpected type %T\n", x)
    }
}
