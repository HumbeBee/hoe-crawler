package definitions

import "time"

type BypassResult struct {
	Success             bool
	IsChallengeDetected bool
	Cookies             []*Cookie
	UserAgent           string
	Error               error
}

type Cookie struct {
	Name    string
	Value   string
	Domain  string
	Path    string
	Secure  bool
	Expires time.Time
}
