package crawler

import (
	"encoding/base64"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/nuronialBlock/crawly/component"
	"github.com/nuronialBlock/crawly/storage"
	log "github.com/sirupsen/logrus"
)

type Crawler struct {
	url string

	component interface{}
}

type crawlerComponent struct {
	frontier *component.Frontier
}

func (c *Crawler) convURLFromString() *url.URL {
	u, err := url.Parse(c.url)
	if err != nil {
		panic(err)
	}
	return u
}

func (c *Crawler) validateURL() {
	u := c.convURLFromString()
	err := component.DNSResolver(u.Host)
	if err != nil {
		log.Error("Failed to resolve dns.")
		panic(err)
	}
}

func (c *Crawler) downloadHTMlContent(url string) (string, error) {
	filename := "temp_" + base64.StdEncoding.EncodeToString([]byte(url))

	err := component.Downloader(url, filename)
	if err != nil {
		log.Error("Failed to download content for, URL: ", url, " error: ", err)
		return "", err
	}

	log.Info("Downloading completed for: ", url)
	return filename, nil
}

func (c *Crawler) contentParser(fileName, url string) []string {
	urls := component.ParseLineAndExtract(fileName, url)
	if urls == nil {
		log.Info("No more URL found to download")
	}

	return urls
}

func (c *Crawler) printLinks(parentLink string, internal, external []string) {
	log.Info("Printing Internal Links for : ", parentLink)
	for _, u := range internal {
		log.Info(u)
	}

	log.Info("Printing External Links for : ", parentLink)
	for _, u := range external {
		log.Info(u)
	}
}

func (c *Crawler) filterLink(urls []string) (internal, external []string) {
	return component.URLFilter(urls, c.convURLFromString().Host)
}

func (c *Crawler) pollURL() string {
	return c.frontierStorage().GetAndDelURLToBeDownloaded()
}

func (c *Crawler) crawl(url string) {
	log.Info("Started Crawling for URL:", url)

	file, err := c.downloadHTMlContent(url)
	if err != nil {
		log.Info("Download completed.")
		return
	}

	u := c.contentParser(file, c.url)
	if u == nil {
		return
	}

	internal, external := c.filterLink(u)

	c.printLinks(url, internal, external)

	// Delete from In Progress and move to Downloaded
	c.frontierStorage().DeleteFromInProgress(url)
	c.frontierStorage().PutDownloaded(url)

	// Push all the new links to be downloaded
	for _, u := range internal {
		if c.frontierStorage().GetDownloaded(u) {
			continue
		}
		c.frontierStorage().PutToBeDownloaded(u)
	}

	// Delete the file created
	m := sync.Mutex{}
	m.Lock()
	err = os.Remove(file)
	if err != nil {
		log.Fatal("Error deleting file...")
		panic(err)
	}
	m.Unlock()
	log.Info("Deleted file: ", file)

	return
}

func (c *Crawler) frontier() *component.Frontier {
	comp := c.component.(crawlerComponent)
	return comp.frontier
}

func (c *Crawler) frontierStorage() *storage.FrontierStorage {
	comp := c.component.(crawlerComponent)
	fs := comp.frontier.FrontierStorage
	return fs
}
func (c *Crawler) Cleanup() {
	log.Info("TODO:// Cleanup resources before closing")
}

func (c *Crawler) init() {
	c.validateURL()
	c.frontierStorage().PutToBeDownloaded(c.url)
}

// Start starts goroutines to crawl based on politeness.
func (c *Crawler) Start() {
	f := component.NewFrontier()
	c.component = crawlerComponent{f}
	c.init()

	for {
		noTask := 0
		numberOfWorker := c.frontier().Politeness

		ch := make(chan int, numberOfWorker)
		for i := 0; i < numberOfWorker; i++ {
			m := sync.Mutex{}
			m.Lock()

			u := c.pollURL()
			if u == "" {
				noTask += 1
				log.Info("No more URL to crawl...")
				m.Unlock()
				continue
			}

			m.Unlock()

			log.Info("Starting for worker - ", i, " URL: ", u)
			go func() {
				// put in progress and start download
				c.frontierStorage().PutInProgress(u)
				c.crawl(u)

				ch <- 1
			}()
		}

		<-ch
		log.Info("Sleeping for delay...")
		time.Sleep(time.Second * time.Duration(c.frontier().Delay))

		if noTask == numberOfWorker {
			break
		}
	}
}

func NewCrawler(url string) *Crawler {
	c := &Crawler{
		url: url,
	}
	return c
}
