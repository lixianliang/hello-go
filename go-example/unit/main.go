package main

import (
    "fmt"
    "unit"
    "net/http"
    "log"
)

func handler(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello 123")
}

func main() {
    http.HandleFunc("/", handler)
    //log.Fatal(http.ListenAndServe(":8500", nil))
    log.Fatal(unit.ListenAndServe(":8400", nil))
}
