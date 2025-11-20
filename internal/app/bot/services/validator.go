package services

import (
	"net/url"
	"strings"
)

func IsValidURL(message string) bool {
	messageTrim := strings.TrimSpace(message)
	if messageTrim == "" {
		return false
	}

	urlParced, err := url.ParseRequestURI(messageTrim)
	if err != nil {
		return false
	}

	if urlParced.Scheme == "" || urlParced.Host == "" {
		return false
	}

	switch strings.ToLower(urlParced.Scheme) {
	case "http", "https":
		return true
	default:
		return false
	}
}
