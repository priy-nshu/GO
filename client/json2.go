package client

import (
	"encoding/json"
)

func marshalAll[T any](items []T) ([][]byte, error) {
    var ret [][]byte
    for _,item:=range items{
    r,err:=json.Marshal(item)
    if err !=nil{
     return nil,err
    }
    ret=append(ret,r)
    }
    return ret,nil
}
