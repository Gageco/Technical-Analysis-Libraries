package main

import (
  "testLibs/RSI"
  "fmt"
  "math/rand"
  // "errors"
)

func main() {

  s := rsi.RSIndex{}
  var err error

  // fmt.Println(rsi.Test())
  // fmt.Println(rand.Float64())

  for i:=0;i<15;i++ {
    s, _ = s.AddPeriod(float64(-i) + rand.Float64())
  }
  if err != nil {
    fmt.Println(err)
  }
  // fmt.Println(s)
  s, _ = s.CalculateRSI()

  fmt.Println(s)
}
