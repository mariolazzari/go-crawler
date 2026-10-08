package main

import (
	"fmt"
	"net/url"
)

func normalizeURL(URL string) (string, error) {

	parsed, err := url.Parse(URL)
	if err != nil {
		return "", err
	}

	nomalized := fmt.Sprintf("%s%s", parsed.Host, parsed.Path)

	return nomalized, nil
}
