package main

import (
	"fmt"
	"log"
    "encoding/json"
)

var data = `
{
    "a": "Easy!",
    "b": {
        "c": 2,
        "dd": [3, 4]
    }
}
`

type T struct {
	A string
	B struct {
		RenamedC int   `json:"c"`
		Dd       []int `json:",flow"`
	}
}

func main() {
	t := T{}

	err := json.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)

	d, err := json.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	m := make(map[interface{}]interface{})

	err = json.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)

	d, err = json.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))
}
