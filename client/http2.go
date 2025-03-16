package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getIssueData(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	// ?
}

