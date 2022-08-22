package cmd

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/nuronialBlock/crawly/crawler"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var url string

// crawlCmd represents the crawl command
var crawlCmd = &cobra.Command{
	Use:              "crawl",
	TraverseChildren: true,
	Short:            "crawl starts crawling the given website domain",
	Run:              run,
}

func handleCleanup(c *crawler.Crawler) {
	c.Cleanup()
}

func run(cmd *cobra.Command, args []string) {
	logrus.Info("Starting crawler for: ", url)

	gracefulShutdown := make(chan os.Signal)
	signal.Notify(gracefulShutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGIO)

	go func() {
		s := <-gracefulShutdown
		log.Info("Starting Cleanup", s)
		os.Exit(1)
	}()

	c := crawler.NewCrawler(url)
	c.Start()
}

func init() {
	rootCmd.AddCommand(crawlCmd)
	crawlCmd.Flags().StringVarP(&url, "URL", "u", "", "URL takes input for the crawling URL")
}
