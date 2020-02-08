package weather

import (
	"errors"
	. "github.com/stretchr/testify/assert"
	"testing"
)

type FailingService struct{}

func (s *FailingService) Report(query string) (*Report, error) {
	return nil, errors.New("some error")
}

func (s *FailingService) Name() string {
	return "FailingService"
}

type SuccessfulService struct{}

func (s *SuccessfulService) Report(query string) (*Report, error) {
	return &Report{WindSpeed: 10, Temperature: 100}, nil
}

func (s *SuccessfulService) Name() string {
	return "SuccessfulService"
}

func TestFailover(t *testing.T) {
	ws := WService{
		Services: []Service{
			new(FailingService),
			new(SuccessfulService),
		},
	}

	report, err := ws.Report("somequery")

	Nil(t, err, "service.Report should not fail")
	Equal(t, report.Temperature, 100, "Temperature mismatch")
}
