package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func peopleApi(sendCh chan int, recCh chan string) {
	for i := range sendCh {
		response, err := http.Get("https://swapi.dev/api/people/" + strconv.Itoa(i))

		if err != nil {
			fmt.Print(err.Error())
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		recCh <- string(responseData)
		fmt.Println("\n")
	}
	defer close(sendCh)
	defer close(recCh)

}
func main() {
	sendCh := make(chan int, 10)
	recCh := make(chan string, 10)
	for i := 1; i <= 10; i++ {
		go peopleApi(sendCh, recCh)
		sendCh <- i
	}

	for resp := range recCh {
		fmt.Print(resp)
	}
	fmt.Println("Printed api responses for 1 to 10 people")
}
