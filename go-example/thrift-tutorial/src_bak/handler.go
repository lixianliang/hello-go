// +build !go1.7

package main

import (
    "fmt"
    "shared"
    "strconv"
    "tutorial"
)

type CalculatorHandler struct {
    log map[int]*shared.SharedStruct
}

func NewCallculatorHandler() *CalculatorHandler {
    return &CalculatorHandler{log: make(map[int]*shared.SharedStruct)}
}

func (p *CalculatorHandler) Add(num1 int32, num2 int32) (retval17 int32, err error) {
    fmt.Print("add(", num1, ",", num2, ")\n")
    return num1 + num2, nil
}

func (p *CalculatorHandler) Calculate(logid int32, w *tutorial.Work) (val int32, err error) {
    fmt.Print("calculate(", logid, ", {", w.Op, ",", w.Num1, ",", w.Num2, "})\n")
    switch w.Op {
    case tutorial.Operation_ADD:
        val = w.Num1 + w.Num2
        break
    case tutorial.Operation_SUBTRACT:
        val = w.Num1 - w.Num2
        break
    case tutorial.Operation_MULTIPLY:
        val = w.Num1 * w.Num2
        break
    case tutorial.Operation_DIVIDE:
        if w.Num2 == 0 {
            ouch := tutorial.NewInvalidOpertion()
            ouch.WhatOp = int32(w.Op)
            ouch.Why = "Cannot divide by 0"
            err = ouch
            return
        }
        val = w.Num1 / w.Num2
        break
    default:
        ouch := tutorial.NewInvalidOperation()
        ouch.WhatOp = int32(w.Op)
        ouch.Why = "Unknown opertation"
        err = ouch
        return
    }
    
    entry := shared.NewSharedStruct()
    entry.Key = logid
    entry.Value = strconv.Itoa(int(val))
    k := int(logid)
    p.log[k] = entry
    return val, err
}

func (p *CalculatorHandler) GetStruct(key int32) (*shared.SharedStruct, error) {
    fmt.Print("getStruct(", key, ")\n")
    v, _ := p.log[int(key)]
    return v, nil
}

func (p *CalculatorHandler) Zip() (err error) {
    fmt.Print("zip()\n")
    return nil
}
