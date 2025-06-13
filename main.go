package main

import (
	"log"
	"go_programs/client"
)

func main() {
	url := "https://api.boot.dev/v1/courses_rest_api/learn-http/users"
	users, err := client.GetUsers(url)
	if err != nil {
		log.Fatal("Error getting users:", err)
	}
	client.LogUsers(users)
}
