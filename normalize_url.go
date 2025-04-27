package main

import (
	"fmt"
	"net/url"
)

func normalizeURL(original_url string) (string, error) {
	parsedUrl, err := url.Parse(original_url)
	if err != nil {
		return "", err
	}
	normalizedUrl := parsedUrl.Hostname() + parsedUrl.EscapedPath()
	fmt.Println(normalizedUrl)
	return normalizedUrl, nil
}


