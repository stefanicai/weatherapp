package main

import (
	"encoding/json"
	"fmt"
	"github.com/stefanicai/weatherapp/weather/openweather"
	"github.com/stefanicai/weatherapp/weather/weatherstack"
	"log"
	"net/http"

	"github.com/stefanicai/weatherapp/weather"
)

func main() {
	weatherService := weather.WService{
		Services: []weather.Service{
			weatherstack.Service{},
			openweather.Service{},
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		query := "Melbourne,AU" //default to Melbourne.

		//parse for query param if provided.
		queries, ok := r.URL.Query()["query"]
		if ok && len(queries[0]) > 0 {
			query = queries[0]
		}

		//get weather report
		report, err := weatherService.Report(query)

		//respond to client
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			//return report as json
			b, err := json.Marshal(report)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			fmt.Fprintf(w, "%s", string(b))
		}
	})

	//start service and log if that fails
	//TODO: add support cross origin - not needed for this test, but will be needed if this is
	// to be provided as a service for others to embed it in their websites/apps.
	log.Println("Starting server at http://localhost:8080/")
	log.Println("E.g. Open http://localhost:8080/query=Melbourne,AU to see weather in Melbourne")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
