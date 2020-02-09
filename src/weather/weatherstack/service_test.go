package weatherstack

import (
	. "github.com/stretchr/testify/assert"
	"testing"
)

func TestConversionValidJson(t *testing.T) {
	s := []byte(`{"current":{"wind_speed":10.1, "temperature":101.3}}`)
	report, err := extract(s)
	Nil(t, err)
	NotNil(t, report)
	Equal(t, 101.3, report.Temperature)
	Equal(t, 10.1, report.WindSpeed)
}

func TestConversionInvalidJson(t *testing.T) {
	s := []byte(`{"current":{"wind_speed":"not float", "temperature":"101.3"}}`)
	report, err := extract(s)
	NotNil(t, err)
	Nil(t, report)
}

func TestConversionInvalidJsonMissingFields(t *testing.T) {
	s := []byte(`{"current":{"wind_speed":10.1}`)
	report, err := extract(s)
	NotNil(t, err)
	Nil(t, report)
}
