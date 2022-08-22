package component

import (
	"io"
	"net/http"
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
)

// Downloader will do dns resolve
// download the content
// and return the downloaded content
func Downloader(url, filename string) (err error) {
	log.Info("Start downloading... ", url)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	m := sync.Mutex{}

	m.Lock()
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	m.Unlock()

	return nil
}
