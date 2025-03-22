package client

type ParsedURL struct {
	Protocol string
	Username string
	Password string
	Hostname string
	Port     string
	Pathname string
	Search   string
	Hash     string
}

