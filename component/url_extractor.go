package component

import (
	"net/url"
	"regexp"

	log "github.com/sirupsen/logrus"
)

var linkMatcher = regexp.MustCompile(".*?<a.*?href=\"(.*?)\"")

func urlExtractor(content string) (*url.URL, error) {
	match := linkMatcher.FindAllStringSubmatch(content, -1)
	if match == nil || len(match) <= 0 {
		return nil, nil
	}

	l, err := url.Parse(match[0][1])
	if err != nil {
		log.Error("URL parsing failed, Error: ", err)
		return nil, err
	}

	return l, nil
}
