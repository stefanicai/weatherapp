package weatherstack

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stefanicai/weatherapp/weather"
	"io/ioutil"
	"net/http"
)

//NOTE: the access key is not valid, thus this won't work.
//That is intentional so this service call fails. Pls replace with correct key if you'd like to make it work.
//I haven't tested the implementation with a valid key, but it should work fine.
const (
	//AccessKey identifies the caller of the API
	AccessKey = "YOUR_ACCESS_KEY" //invalid key

	//CallURL is a template used to build the full Get URL required to call the server
	CallURL = "http://api.weatherstack.com/current?access_key=%s/query=%s"
)

//Service is the implementation of open weather for weather.Service.
type Service struct {
}

//Report the weather
func (s Service) Report(query string) (*weather.Report, error) {
	url := fmt.Sprintf(CallURL, AccessKey, query)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	respData, _ := ioutil.ReadAll(resp.Body)
	report, err := extract(respData)
	return report, err
}

//Name of the service
func (s Service) Name() string {
	return "WeatherStack"
}

//extract parses the provided json and extracts required values
//It assumes that the json is in the correct format, as per the API specs
func extract(j []byte) (r *weather.Report, e error) {
	result := jsonResult{}
	e = json.Unmarshal(j, &result)
	if e == nil {
		if result.Success == false {
			//get error
			e = errors.New(fmt.Sprintf("Service error - type: %s - info: %s", result.Err.ErrType, result.Err.Info))
		} else {
			r = new(weather.Report)
			r.WindSpeed = result.Current.WindSpeed
			r.Temperature = result.Current.Temperature
		}
	}
	return r, e
}

//create JSON structures to extract values. Works well for simple structures, but can be a bit much for more complex
//json-s where we only need some fields. Using this here as a different method.
type jsonReport struct {
	WindSpeed   float64 `json:"wind_speed"`
	Temperature float64 `json:"temperature"`
}

type jsonResult struct {
	Current jsonReport
	Success bool
	Err     CallError `json:"error"`
}

type CallError struct {
	Code    int
	ErrType string `json:"type"`
	Info    string
}
