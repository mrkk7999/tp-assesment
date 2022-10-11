package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func peopleApi(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	response, err := http.Get("https://swapi.dev/api/people/" + strconv.Itoa(i))
	if err != nil {
		fmt.Print(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
	fmt.Println("\n")
}
func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		// go API(myChan, respChan)
		// myChan <- i
		wg.Add(1)
		go peopleApi(i, &wg)
	}
	wg.Wait()
	fmt.Println("Printed api responses for 1 to 10 people")
}
