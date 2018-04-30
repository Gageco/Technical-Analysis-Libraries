package rsi
///*****************************************************************************
// little bit of background on RSI, it stands for Relative Strength Index
// it measures the speed and change of price movement. invented by
// J. Welles Wilder around 1978. Typically for a stock an RSI over 70 is
// considered overbought and under 30 is undersold. You can read more about it
// here http://stockcharts.com/school/doku.php?id=chart_school:technical_indicators:relative_strength_index_rsi
//******************************************************************************

//******************************************************************************
// NOTE for future Gage, see if you can set what is considered overbought and
// over sold, probably just modify the struct
//******************************************************************************
import (
  "fmt"
  "errors"
  "math"
)

func Test() string {
  fmt.Println("Test From Library")
  return "hello, world"
}

func (s RSIndex) AddLength(l int) (RSIndex) {                                   // add a length to RSIndex
  s.Length = l
  return s
}

func (s RSIndex) AddPeriod(p float64) (RSIndex, error) {                          // add a period to RSIndex
  //****************************************************************************
  // for RSI the default length is 0 and if the library user does not put in
  // a length before then it defaults to 15
  //****************************************************************************
  if s.Length == 0 {
    s.Length = 14
  }

  //****************************************************************************
  // if the price slice is not full than it will add a new one to the end of the
  // slice
  //****************************************************************************
  if len(s.Periods) < s.Length {
    s.Periods = append(s.Periods, p)
    return s, nil
  }

  //****************************************************************************
  // if the slice is full than it will remove the first one, move all down one
  // and add one to the end
  //****************************************************************************
  s.Periods = append(s.Periods[:0], s.Periods[1:]...)
  s.Periods = append(s.Periods, p)
  return s, nil

}

func (s RSIndex) CalculateRSI() (RSIndex, error) {
  //****************************************************************************
  // to calculate RSI you need the array to be of sufficient length otherwise
  // it cannot be calculated currectly
  //****************************************************************************
  if s.Length != len(s.Periods) {
    err := errors.New("RSI cannot be calcuated because the length is not sufficient, add more data points using .AddPeriod(int)")
    return s, err
  }

  var positive []float64
  var avgPositive float64
  var sumPositive float64

  var negative []float64
  var avgNegative float64
  var sumNegative float64

  //****************************************************************************
  // for all items in array calculate the change in price and then add it to the
  // positive array or the negative array. Then calcualte the RS then RSI
  //****************************************************************************
  for i:=0; i<s.Length-1;i++ {
    delta := s.Periods[i+1] - s.Periods[i]
    if delta >= 0 {
      positive = append(positive, delta)
    } else {
      negative = append(negative, delta)
    }
  }

  //****************************************************************************
  // getting positive average
  //****************************************************************************
  for i:=0; i < len(positive);i++ {
    sumPositive += positive[i]
  }
  avgPositive = sumPositive/float64(len(positive))
  fmt.Println(avgPositive)

  //****************************************************************************
  // getting negative average
  //****************************************************************************
  for i:=0;i<len(negative);i++ {
    sumNegative += math.Abs(negative[i])
  }
  avgNegative = sumNegative/float64(len(negative))
  fmt.Println(avgNegative)

  //****************************************************************************
  // calculating RS, RSI
  //****************************************************************************

  var rs float64
  var rsi float64

  if math.IsNaN(avgNegative) {
    rsi = float64(100)
  } else if math.IsNaN(avgPositive) {
    rsi = float64(0)
  } else {
    rs = avgPositive/avgNegative
    rsi = float64(100) - float64(100)/(float64(1) + rs)
  }

  s.RS = rs
  s.RSI = rsi

  if rsi >= float64(70) {
    s.State = "overbought"
  } else if rsi <= float64(30) {
    s.State = "oversold"
  } else {
    s.State = "neutral"
  }

  return s, nil

}
