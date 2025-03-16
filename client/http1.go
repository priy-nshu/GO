package client

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func getIssues(url string) ([]Issue, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	// ?
	var issue []Issue
	decode :=json.NewDecoder(res.Body)
	if err:=decode.Decode(&issue);err!=nil{
		fmt.Println("error decoding response body")
		return nil,err
	}
	return issue,nil
}

