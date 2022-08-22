# crawly
Crawly crawls the website link you give and prints the links and visit under the same domain.

# Pre-requesite/Steps for running

- install Golang

- export GOPATH of the working directory

- from commandline run -  `go install`

- to run the application, run - `go run main.go crawl -u http://yoururl.com`

- By default this app sets the concurrency to crawl(politeness) 1, and 60 second delay between two consecutive url crawl. If you want to change this export politeness and delay to desired number.


### Room of Improvements [In Code]

- Worker pool is not implemented yet to control the flow of concurrency.

- The crawling logic heavily attached to crawl.go file. Some operations can be decoupled. For example, pollinmg from queue can be moved to Frontier.

- For graceful shutdown, cleaning up the resources has not been finished. 

- For future reference, a lot of config variables can be read from Config file

- Complete and keep test coverage more than 95%