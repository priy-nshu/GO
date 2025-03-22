package client
/*
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getIssues(url string) ([]Issue, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data ,err:=io.ReadAll(res.Body)
	if err!=nil{
	 return nil,err
	}
	var issue []Issue
	if err:=json.Unmarshal(data,&issue);err!=nil{
	  return nil,err
	}
	return issue,nil
}
*/
