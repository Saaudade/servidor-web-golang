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

	for _, server := range servers {
		go checkServer(server, channel)
	}

	for i := 0; i < len(servers); i++ {
		fmt.Println(<-channel)
	}

	timeSpent := time.Since(start)
	fmt.Printf("Execution time %s\n", timeSpent)
}

func checkServer(server string, channel chan string) {
	_, err := http.Get(server)

	if err != nil {
		fmt.Println(server, "Not found =(")
		channel <- server + " Not found =("
	} else {
		fmt.Println(server, "It is operating normally =)")
		channel <- server + " It is operating normally =)"
	}
}
