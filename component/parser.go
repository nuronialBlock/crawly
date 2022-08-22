package component

import (
	"bufio"
	"net/url"
	"os"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func trimSuffix(u, suffix string) string {
	if strings.HasSuffix(u, suffix) {
		u = u[:len(u)-len(suffix)]
	}
	return u
}

func urlPurifier(u *url.URL, baseURL string) string {
	if u.IsAbs() {
		return u.String()
	}

	bu, err := url.Parse(baseURL)
	if err != nil {
		log.Error("Failed to parse URL: ", baseURL)
	}

	// case: //bn-in.wikipedia.com
	matcher := regexp.MustCompile("^/[a-z/_-]+[.]+").MatchString(u.String())
	if matcher {
		return bu.Scheme + ":" + u.String()
	}

	return bu.Scheme + "://" + trimSuffix(bu.Host, "/") + u.String()
}

// ParseLineAndExtract parses html content line by line from downlonloaded content
// and returns urls.
func ParseLineAndExtract(filename, baseURL string) []string {
	file, err := os.Open(filename)
	checkErr(err)

	log.Info("Started scanning...")
	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		u, err := urlExtractor(scanner.Text())

		if err != nil {
			return nil
		}

		if u == nil {
			continue
		}

		p := urlPurifier(u, baseURL)
		urls = append(urls, trimSuffix(p, "/"))
	}

	return urls
}
