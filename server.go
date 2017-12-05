package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/syllabix/ecaas"
)

type demoMoveType struct{}

func (mt *demoMoveType) GetName() string {
	return "Demo"
}

func (mt *demoMoveType) GetMultiplier() string {
	return "0.05"
}

func (mt *demoMoveType) GetTaxRate() string {
	return "0.06"
}

func handler(w http.ResponseWriter, r *http.Request) {
	var jobDetails ecaas.JobDetails
	err := json.NewDecoder(r.Body).Decode(&jobDetails)
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, "Error parsing job details")
		return
	}

	moveType := demoMoveType{}
	results := ecaas.CalculateTotalCost(&jobDetails, &moveType)
	fmt.Fprintf(w, "Estimate range is %v - %v", results.Low, results.High)
}

func main() {
	port := os.Getenv("HTTP_PLATFORM_PORT")
	if len(port) == 0 {
		port = "8080"
	}

	http.HandleFunc("/", handler)

	log.Println("Starting webserver on port ", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Println("Error starting webserver: ", err)
	}
}
