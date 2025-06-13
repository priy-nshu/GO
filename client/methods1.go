package client

import (
	"encoding/json"
	"net/http"
	"io"
)

func GetUsers(url string) ([]User, error) {
	// ?
	var usr []User
	resp,err:=http.Get(url)
	
	if err!=nil{
	 return nil,err
	}
	defer resp.Body.Close()
	
	body,err:=io.ReadAll(resp.Body)
	if err !=nil{
	 return nil,err
	}
	
	if err:=json.Unmarshal(body,&usr);err!=nil{
	 return nil,err
	}
	return usr,nil
}
