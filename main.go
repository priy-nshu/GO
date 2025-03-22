package main

/*import(
    "log"
)
*/

import (
    "go_programs/client"
    "net/url"
)
/*
func main() {
	issues, err := client.GetIssues(client.Domain)
	if err != nil {
		log.Fatalf("error getting issues data: %v", err)
	}
	client.LogIssues(issues)
}*/

func NewParsedURL(urlString string) ParsedURL {
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		return ParsedURL{}
	}

	password, _ := parsedUrl.User.Password()

	return ParsedURL{
		Protocol: parsedUrl.Scheme,
		Username: parsedUrl.User.Username(),
		Password: password,
		Hostname: parsedUrl.Hostname(),
		Port:     parsedUrl.Port(),
		Pathname: parsedUrl.Path,
		Search:   parsedUrl.RawQuery,
		Hash:     parsedUrl.Fragment,
	}
}
