package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type webdata struct {
	Number int `json:"price"`
}

func main() {
	fmt.Println("Welcome. This is the Go example without go routines.")
	start := time.Now()

	urls := []string{
		"http://localhost:4501/number1",
		"http://localhost:4501/number2",
		"http://localhost:4501/number3",
	}

	var total int

	for _, web := range urls {

		result := getNumberFromWeb(web)

		total += result

	}

	fmt.Println("Total price is: ", total)

	duration := time.Since(start)
	fmt.Println("Execution time is: ", duration)
}

func getNumberFromWeb(url string) int {

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	return decodeJson(content).Number
}

func decodeJson(dataFromWeb []byte) webdata {

	var received webdata

	checkValid := json.Valid(dataFromWeb)

	if checkValid {
		json.Unmarshal(dataFromWeb, &received)
	}

	return received

}
