package validator

import "net/url"

func IsValidURL(rawURL string) bool {
	u, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return false
	}

	if u.Host == "" {
		return false
	}

	switch u.Scheme {
	case "http", "https":
		return true
	default:
		return false
	}
}
