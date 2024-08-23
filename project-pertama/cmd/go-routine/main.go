package main

import (
	"fmt"
	"io"
	"net/http"
	"runtime"
	"sync"
	"time"
)

func example1() {
	fmt.Println("Hello, wolrd!")

	go sayHello("Hello, Go Routine!")

	sayHello("Hello, Synchronous!")

	go func() {
		fmt.Println("Hello, Anonymous Go Routine!")
	}()

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("Nilai I: ", i)
		}()
	}

	fmt.Println("End of main function")

	fmt.Println("Number of CPUs: ", runtime.NumCPU())
	fmt.Println("Number of Go Routines: ", runtime.NumGoroutine())

	// sleep for 1 second
	// to allow the go routine to finish
	time.Sleep(1 * time.Second)
}

func exampleWaitGroup() {
	start := time.Now()
	fruits := []string{"Apple", "Mango", "Banana", "Pineapple", "Watermelon"}

	var wg sync.WaitGroup

	for _, fruit := range fruits {
		wg.Add(1)

		go func(fruit string) {
			defer wg.Done()

			apiCall()
		}(fruit)
	}

	wg.Wait()

	fmt.Printf("Elapsed time: %v\n", time.Since(start))
}

func apiCall() {

	url := "https://jsonplaceholder.typicode.com/posts/1"

	// call the API
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("API called")

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bodyBytes))
}

func exampleWithoutWaitGroup() {
	start := time.Now()
	fruits := []string{"Apple", "Mango", "Banana", "Pineapple", "Watermelon"}

	for range fruits {
		// go func(fruit string) {
		apiCall()
	}

	fmt.Printf("Elapsed time: %v\n", time.Since(start))
}

func main() {
	exampleWaitGroup()
	// exampleWithoutWaitGroup()

}

func sayHello(message string) {
	fmt.Println(message)
}
