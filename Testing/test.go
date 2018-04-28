package main

import (
  "testLibs/SMA"
  "fmt"
  // "errors"
)

func main() {

  s := sma.Periods{}
  d := sma.Periods{}

  d = d.AddLength(50)
  s = s.AddLength(20)

  for i:=0;i<55;i++ {
    s = s.AddPrice(float64(i))
    d = d.AddPrice(float64(i)*1.5)
  }

  fmt.Println(s.Calculate())
  fmt.Println(d.Calculate())

  f := sma.RelativeSMA{}
  f = f.AddRelativeSMA(s,d)
  fmt.Println(f)

  // fmt.Println(sma.Test())

}
