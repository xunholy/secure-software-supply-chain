package main

import "fmt"

var (
	appName string = "N/A"
	version string = "N/A"
	sha     string = "N/A"
)

func main() {
	fmt.Printf("Application: %v, Version: %v, SHA: %v", appName, version, sha)
}
