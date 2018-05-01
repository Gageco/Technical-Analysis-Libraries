package main

import (
  "testLibs/PoloniexInterface"
  "testLibs/SMA"
  "fmt"
  "time"
  // "strconv"
  // "math/rand"
  // "errors"
)

func main() {
  var err error

  simpleMA := sma.Periods{}
  simpleMA = simpleMA.SetLength(20)

  err = err
  s := polo.HistoricalPoints{}

  t := time.Now()
  t2 := t.Add(-24 * 40 * time.Hour)

  // s, _ = s.SetPeriodLength(86400)

  s, err = s.GetHistoricalData(t2, t)
  if err != nil {
    fmt.Println(err)
  }

  // fmt.Println(s)
  for i:=0;i<len(s.Points);i++ {
    simpleMA = simpleMA.AddPrice(s.Points[i].Close)
    fmt.Println(simpleMA.CalculateSMA())
  }



  // fmt.Println(s)
}
