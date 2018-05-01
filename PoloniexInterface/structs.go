package polo

import (
  "time"
)

type HistoricalPoint struct {
  Date          int64  `json:"date"`
  High          float64    `json:"high"`
  Low           float64    `json:"low"`
  Open          float64    `json:"open"`
  Close         float64    `json:"close"`
  Volume        float64    `json:"volume"`
  QuoteVolume   float64    `json:"quoteVolume"`
  WeightedAvg   float64    `json:"weightedAverage"`
}

type HistoricalPoints struct {
  PeriodLength     int
  Pair             string
  StartTime        time.Time
  EndTime          time.Time
  Points           []HistoricalPoint
}
