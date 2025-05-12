package main

type TwoValues struct {
    Value1 float64 `json:value1`
    Value2 float64 `json:value2`
}

type SingleValue  struct {
    Value1 float64 `json:value1`
}

type ArrayValue struct  {
    Value1 float64[] `json:value1`
}
