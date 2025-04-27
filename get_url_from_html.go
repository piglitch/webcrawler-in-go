package main

import (
	"strings"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getURLsFromHTML(htmlBody, rawbaseURL string) ([]string, error) {
	absoluteUrls := []string{}
	htmlReader := strings.NewReader(htmlBody)
	htmlNode, err := html.Parse(htmlReader)
	if err != nil {
		return []string{}, err
	}
	for n := range htmlNode.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					absurl := a.Val
					if a.Val[0] != 'h' {
						absurl = rawbaseURL + absurl 
					}
					absoluteUrls = append(absoluteUrls, absurl)
				}
			}
		}
	}
	return absoluteUrls, nil
}