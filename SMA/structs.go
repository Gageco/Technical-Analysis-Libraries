package sma

import (
  // "time"
)

type Periods struct {
  Length    int
  Prices   []float64
}

type RelativeSMA struct {
  SMA1           []Periods
  SMA2           []Periods
  Difference     float64
  NearCross      bool
  CrossType      string
}
