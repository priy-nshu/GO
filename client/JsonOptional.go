package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

func getResources(url string) ([]map[string]any, error) {
	var resources []map[string]any

	res, err := http.Get(url)
	if err != nil {
		return resources, err
	}

	defer res.Body.Close()

if err := json.NewDecoder(res.Body).Decode(&resources); err != nil {
		return resources, err
	}

	return resources, nil
}

func logResources(resources []map[string]any) {
	var formattedStrings []string
	for _,resource := range resources {
		for key,value:=range resource{
		formattedStrings = append(formattedStrings, fmt.Sprintf("Key: %s - Value: %v ",key,value))
	}
	}
	sort.Strings(formattedStrings)

	for _, str := range formattedStrings {
		fmt.Println(str)
	}
}

