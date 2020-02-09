package openweather

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/spyzhov/ajson"
	"github.com/stefanicai/weatherapp/weather"
)

const (
	//APIVersion is the version of api to be used
	APIVersion = "2.5"
	//AppID identifies the caller of the API
	AppID = "2326504fb9b100bee21400190e4dbe6d"

	//CallURL is a template used to build the full Get URL required to call the server
	CallURL = "http://api.openweathermap.org/data/%s/weather?appid=%s&q=%s"
)

//Service is the implementation of open weather for weather.Service.
type Service struct {
}

//Report the weather
func (s Service) Report(query string) (*weather.Report, error) {
	url := fmt.Sprintf(CallURL, APIVersion, AppID, query)
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
	return "OpenWeather"
}

//extract parses the provided json and extracts required values
//It assumes that the json is in the correct format, as per the API specs
func extract(j []byte) (r *weather.Report, e error) {
	defer func() {
		if r := recover(); r != nil {
			e = r.(error)
		}
	}()

	node, err := ajson.Unmarshal(j)
	if err != nil {
		return r, err
	}

	//TODO: Same sa for weatherstack, parse in case of error too. See that package for an example.

	r = new(weather.Report)

	//API returns one value only
	values, err := node.JSONPath("$.main.temp")
	temperature, err := strconv.ParseFloat(values[0].String(), 64)
	if err != nil {
		return nil, err
	} else {
		r.Temperature = temperature
	}

	values, err = node.JSONPath("$.wind.speed")
	windSpeed, err := strconv.ParseFloat(values[0].String(), 64)
	if err != nil {
		return nil, err
	} else {
		r.WindSpeed = windSpeed
	}

	return r, nil
}
