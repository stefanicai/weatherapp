package openweather

import (
	. "github.com/stretchr/testify/assert"
	"testing"
)

func TestConversionValidJson(t *testing.T) {
	s := []byte(`{"main":{"temp":10.1},"wind":{"speed":101.3}}`)
	report, err := extract(s)
	Nil(t, err)
	NotNil(t, report)
	Equal(t, 10.1, report.Temperature)
	Equal(t, 101.3, report.WindSpeed)
}

func TestConversionInvalidJson(t *testing.T) {
	s := []byte(`{"main":{"temp":"10.1""},"wind":{"speed":"not number"}}`)
	report, err := extract(s)
	NotNil(t, err)
	Nil(t, report)
}

func TestConversionInvalidJsonMissingFields(t *testing.T) {
	s := []byte(`{"main":{"tempNot":"10.1""},"windNot":{"speed":"not number"}}`)
	report, err := extract(s)
	NotNil(t, err)
	Nil(t, report)
}
