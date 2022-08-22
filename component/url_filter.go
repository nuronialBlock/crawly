package component

import (
	"net/url"
)

var internalLink, externalLink []string

func URLFilter(urls []string, domain string) (internalLink, externalLink []string) {
	isURLSeen := map[string]bool{}
	for _, u := range urls {
		if isURLSeen[u] {
			continue
		}
		isURLSeen[u] = true

		url, _ := url.Parse(u)
		if url.Hostname() == domain {
			internalLink = append(internalLink, u)
			continue
		}
		externalLink = append(externalLink, u)
	}

	return internalLink, externalLink
}
