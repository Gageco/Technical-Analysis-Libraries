package sma

import (
  "fmt"
  "math"
  // "errors"
  // "time"
)

func Test() string {
  fmt.Println("Test From Library")
  return "hello, world"
}

func (s Periods) Calculate() float64 {                                          //Calculates the average of all points in Period
  total := float64(0.00)
  for i:=0; i<len(s.Prices); i++ {
    total += float64(s.Prices[i])
  }
  average := total/float64(s.Length)
  return average
}

func (s Periods) AddLength(len int) (Periods) {                                 //Add a length to Period, or just use .Length = on Period var
  s.Length = len
  return s
}

func (s Periods) AddPrice(price float64) (Periods) {                            //Add new price to Period, remove old one if full
  //****************************************************************************
  // Check if length is 0, commented out due to a 'fun' little Go quirk where
  // if you use := it will always reset the variable and if you want to have
  // errors returned its best to use := otherwise you will have to declare a
  // error varialbe which nobody really does. Figureing this is the only
  // error i accounted for i figured it was easier to comment it out than
  // figure away around this with points and pass by value nonsense
  //****************************************************************************
  // if s.Length == 0 {
  //   err := errors.New("You need to add a length first, use '.AddLength(length)'")
  //   return s, err
  // }

  //****************************************************************************
  // if the price slice is not full than it will add a new one to the end of the
  // slice
  //****************************************************************************
  if len(s.Prices) < s.Length {
    s.Prices = append(s.Prices, price)
    return s
  }

  //****************************************************************************
  // if the slice is full than it will remove the first one, move all down one
  // and add one to the end
  //****************************************************************************
  s.Prices = append(s.Prices[:0], s.Prices[1:]...)
  s.Prices = append(s.Prices, price)
  return s
}

func (rs RelativeSMA) AddRelativeSMA(p1 Periods, p2 Periods) (RelativeSMA) {                  //Add to a RelativeSMA struct and check for crosses
  rs.SMA1 = append(rs.SMA1, p1)
  rs.SMA2 = append(rs.SMA2, p2)
  rs.Difference = p1.Calculate() - p2.Calculate()


  //****************************************************************************
  // (p1 - p2)/p2 will give a percentage difference, if this difference is less
  // than 5% than we can say it is approaching a cross. 5% is arbitrary
  //****************************************************************************
  if math.Abs(rs.Difference/p2.Calculate()) <= .05 {
    rs.NearCross = true
  } else {
    rs.NearCross = false
  }

  //****************************************************************************
  // an attempt to figure out what sort of cross is being approached. I
  // believe it works correct and the crosses it ID's are the right ones
  //****************************************************************************
  if rs.NearCross {
    if rs.Difference > 0 {                                                      // p1 > p2, difference is positive

  //****************************************************************************
  // with a positive difference and p1.Length > p2.Length that means that the
  // approaching cross is going to be a golden cross because you are in bearish
  // state already. 50 SMA is above 20 SMA, the approaching cross is golden
  //****************************************************************************
      if p1.Length > p2.Length {
        rs.CrossType = "golden"

  //****************************************************************************
  // with a positive difference and p1.Length < p2.Length that means that the
  // approaching corss is going to be a death cross because you are in a bullish
  // state already 20 SMA is above 50 SMA, the approachiing cross is death
  //****************************************************************************
      } else {
        rs.CrossType = "death"
      }

    } else if rs.Difference < 0 {                                               // p2 > p1, difference is negative

  //****************************************************************************
  // With a negative distance and p1.Length > p2.Length that means that the
  // approaching cross is going to be a death cross because you are in bullish
  // state already 20 SMA is above 50 SMA, the approaching cross is death
  //****************************************************************************
      if p1.Length > p2.Length {
        rs.CrossType = "death"

  //****************************************************************************
  // With a negative distance and p2.Length > p1.Length that means that the
  // approaching cross is going to be a golden cross because you are in bearish
  // state already 50 SMA is above 20 SMA, the approaching cross is golden
  //****************************************************************************
      } else {
        rs.CrossType = "golden"
      }
    }
  } else {
    rs.CrossType = "none"
  }

  return rs
}
