package sma

import (
  // "time"
)

type Periods struct {
  Length    int
  Prices   []float64
}

type RelativeSMA struct {
  SMA1               []Periods
  SMA2               []Periods
  // CrossDifference    float64
  Difference         float64
  NearCross          bool
  CrossType          string
}
