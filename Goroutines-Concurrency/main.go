package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	channel := make(chan string)

	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	i := 0

	for {
		if i > 3 {
			break
		}
		for _, server := range servers {
			go checkServer(server, channel)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-channel)
		i++
	}

	timeSpent := time.Since(start)
	fmt.Printf("Execution time %s\n", timeSpent)
}

func checkServer(server string, channel chan string) {
	_, err := http.Get(server)

	if err != nil {
		channel <- server + " Not found =("
	} else {
		channel <- server + " It is operating normally =)"
	}
}
