package polo

import (
  "strconv"
  "errors"
  "bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//******************************************************************************
// get historical data from poloniex.com between a certain period of time. Takes
// a start and an end date
//******************************************************************************
func (s HistoricalPoints) GetHistoricalData(start time.Time, end time.Time) (HistoricalPoints, error) {
  var baseURL string = "https://poloniex.com/public?command=returnChartData&"
  var requestURL string

  if s.PeriodLength == 0 {  //Default to a period of 1 day
    s.PeriodLength = 86400
  }
  if s.Pair == "" {         //Default to USDT_BTC as pair
    s.Pair = "USDT_BTC"
  }
  s.StartTime = start
  s.EndTime = end

  //****************************************************************************
  // check that the period length is going to be valid
  //****************************************************************************
  if s.PeriodLength != 300 && s.PeriodLength != 900 && s.PeriodLength != 1800 && s.PeriodLength != 7200 && s.PeriodLength != 14400 && s.PeriodLength != 86400 {
    err := errors.New("Your period length is invalid. Must be either of the following; 300,900,1800,7200,14400,86400")
    return s, err
  }

  //****************************************************************************
  // assemble using variables the request url that is going to be sent to polo
  //****************************************************************************
  requestURL = baseURL + "currencyPair=" + s.Pair + "&start=" + strconv.FormatInt(int64(start.Unix()), 10) + "&end=" + strconv.FormatInt(int64(end.Unix()), 10) + "&period=" + strconv.Itoa(s.PeriodLength)

  //****************************************************************************
  // Do a GET request with this URL and parse it into HistoricalPoint then put
  // that whole thing into HistoricalPoints
  //****************************************************************************
  response, err := http.Get(requestURL)
	if err != nil {
		return s, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return s, err
	}
	data := bytes.TrimSpace(body)
	data = bytes.TrimPrefix(data, []byte("// "))

	err = json.Unmarshal(data, &s.Points)

	if err != nil {
		return s, err
	}

  return s, nil

}

func GetValidPeriodLengths() (string) {
 return "300 (5min), 900(15min), 1800(30min), 7200(2hr), 14400(4hr), 86400(1day))"
}

func (s HistoricalPoints) SetPeriodLength(l int) (HistoricalPoints, error) {

  if l != 300 && s.PeriodLength != 900 && s.PeriodLength != 1800 && s.PeriodLength != 7200 && s.PeriodLength != 14400 && s.PeriodLength != 86400 {
    err := errors.New("Your period length is invalid. Must be either of the following; 300,900,1800,7200,14400,86400")
    return s, err
  }

  s.PeriodLength = l
  return s, nil
}

func (s HistoricalPoints) SetPair(p string) (HistoricalPoints) {
  s.Pair = p
  return s
}

func Test() string {
  fmt.Println("Test From Library")
  return "hello, world"
}
