package client

import (
	"net/url"
)

func getDomainNameFromURL(rawURL string) (string, error) {
    Url,err:=url.Parse(rawURL)
    if err!=nil{
      return "",err
    }
    return Url.Hostname(),nil
}
