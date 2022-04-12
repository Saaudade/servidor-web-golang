package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for _, server := range servers {
		checkServer(server)
	}

	timeSpent := time.Since(start)
	fmt.Printf("Execution time %s\n", timeSpent)
}

func checkServer(server string) {
	_, err := http.Get(server)

	if err != nil {
		fmt.Println(server, "Not found =(")
	} else {
		fmt.Println(server, "It is operating normally =)")
	}
}
