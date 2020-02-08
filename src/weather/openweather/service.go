package openweather

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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
	log.Println(string(respData))
	return &weather.Report{}, nil
}

//Name of the service
func (s Service) Name() string {
	return "OpenWeather"
}
