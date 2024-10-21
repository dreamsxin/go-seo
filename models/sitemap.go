package models

type SitemapConfig struct {
	Dsturl         string
	Savepath       string
	Filename       string
	Concurrency    int
	Crawltimeout   int
	Requesttimeout int
}
