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
	fmt.Println("Welcome. This is the Go example with go routines.")
	start := time.Now()

	urls := []string{
		"http://localhost:4501/number1",
		"http://localhost:4501/number2",
		"http://localhost:4501/number3",
	}

	var total int

	ch := make(chan int)

	for _, web := range urls {

		go getNumberFromWeb(web, ch)

	}

	a, b, c := <-ch, <-ch, <-ch

	total = a + b + c

	fmt.Println("A Value is: ", a)
	fmt.Println("B Value is: ", b)
	fmt.Println("C Value is: ", c)

	fmt.Println("Total price is: ", total)

	duration := time.Since(start)
	fmt.Println("Execution time is: ", duration)
}

func getNumberFromWeb(url string, c chan int) {

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	c <- decodeJson(content).Number
}

func decodeJson(dataFromWeb []byte) webdata {

	var received webdata

	checkValid := json.Valid(dataFromWeb)

	if checkValid {
		json.Unmarshal(dataFromWeb, &received)
	}

	return received

}
