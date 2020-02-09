package weather

import (
	"errors"
	"log"
)

//Report is a weather report
type Report struct {
	WindSpeed   float64
	Temperature float64
}

//Service defines a service that can be called to provide weather info.
type Service interface {
	//Generate report
	Report(query string) (*Report, error)

	//Service name
	Name() string
}

//WService provides a weather report by calling a list of services subsequently until one responds.
type WService struct {
	Services []Service
}

//Report generates a weather report by calling a list of external services subsequently until a valid response is received.
//It returns the report if at least one service responded, or an error if all services failed to respond.
func (ws *WService) Report(query string) (*Report, error) {
	for _, s := range ws.Services {
		report, err := s.Report(query)
		if err == nil {
			return report, nil
		} else {
			//log and try next service.
			log.Printf("Service %s failed with error: %v\n", s.Name(), err)
		}
	}
	return nil, errors.New("Service Unavailable")
}
