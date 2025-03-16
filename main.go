package main

import (
	"fmt"
	"log"
    "go_programs/client"
)

const issueURL = "https://api.boot.dev/v1/courses_rest_api/learn-http/issues"

func main() {
	issues, err := client.GetIssueData(issueURL)
	if err != nil {
		log.Fatalf("error getting issue data: %v", err)
	}
	prettyData, err := client.Prettify(string(issues))
	if err != nil {
		log.Fatalf("error prettifying data: %v", err)
	}
	fmt.Println(prettyData)
}

