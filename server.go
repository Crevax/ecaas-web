package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/syllabix/ecaas"
)

type demoMoveType struct{}

func (self *demoMoveType) GetName() string {
	return "Demo"
}

func (self *demoMoveType) GetMultiplier() string {
	return "0.05"
}

func (self *demoMoveType) GetTaxRate() string {
	return "0.06"
}

func handler(w http.ResponseWriter, r *http.Request) {
	jobDetails := ecaas.NewJobDetails(4, "120.50", "Mon Dec 04 12:00:00 EST 2017")
	moveType := demoMoveType{}
	results := ecaas.CalculateTotalCost(jobDetails, &moveType)
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
