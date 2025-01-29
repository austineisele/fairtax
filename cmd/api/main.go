package main

import (
	"encoding/json"
	"fmt"
	"github.com/austineisele/fairtax/config"
	"github.com/austineisele/fairtax/internal/domain"
	"io"
	"log"
	"net/http"
	"os"
)

type Response struct {
	LevyData []domain.LevyData `json:"levy_entries"`
}

func main() {
	url := config.GetApiUrl("RateLevy")
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(len(responseObject.LevyData))

  i := 1
	for _, levy_data := range responseObject.LevyData {
    fmt.Printf("Entry %d", i)
		fmt.Printf("School Name: %s\n", levy_data.SchoolName)
    fmt.Printf("Municipality: %s\n", levy_data.Municipality)
    fmt.Printf("Municipality: %d\n", levy_data.RollYear)
	}

}
