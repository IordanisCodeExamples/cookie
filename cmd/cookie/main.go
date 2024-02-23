package main

import (
	"flag"
	"fmt"

	"cookie/internal/input"
	"cookie/internal/service"
)

func main() {

	input := input.New()

	srv, err := service.New(input)
	if err != nil {
		panic(fmt.Sprintf("error creating service object %s", err.Error()))
	}

	// Parse command-line arguments
	var fileName, dateStr string
	flag.StringVar(&fileName, "f", "", "Filename of the cookie log")
	flag.StringVar(&dateStr, "d", "", "Date to find the most active cookie")
	flag.Parse()

	// Validate input
	if fileName == "" || dateStr == "" {
		fmt.Println("Both -f and -d parameters are required")
		return
	}

	mostActiveCookies, err := srv.FindMostActiveCookiesForDate(fileName, dateStr)
	if err != nil {
		panic(fmt.Sprintf("error finding most active cookies %s", err.Error()))
	}

	// Print the most active cookies
	for _, cookie := range mostActiveCookies {
		fmt.Println(cookie)
	}
}
