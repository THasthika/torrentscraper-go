package torrentscraper

import (
	"fmt"

	"github.com/tharindu96/torrentscraper-go/scraper"
	"github.com/tharindu96/torrentscraper-go/scraper/eztvag"
	"github.com/tharindu96/torrentscraper-go/scraper/zooqlecom"
)

var scrapers map[string]*scraper.Scraper

// Init func
func Init() {
	scrapers = make(map[string]*scraper.Scraper)

	registerScraper(eztvag.Init())
	registerScraper(zooqlecom.Init())
}

// Search func
func Search(query string) {
	res := searchType(query, scraper.TorrentTypeUnspecified)
	fmt.Println(res)
}

// SearchMovie func
func SearchMovie(query string) {
	res := searchType(query, scraper.TorrentTypeMovie)
	fmt.Println(res)
}

func searchType(query string, ttype scraper.TorrentType) []*scraper.TorrentMeta {
	out := make(chan scraper.Result)

	count := 0
	for _, s := range scrapers {
		if s.SupportedTypes&ttype == ttype {
			go s.Search(query, ttype, out)
			count++
		}
	}

	res := make([]*scraper.TorrentMeta, 0)
	for i := 0; i < count; i++ {
		result := <-out
		fmt.Println(result)
		if result.Err == nil && result.Torrents != nil {
			res = append(res, result.Torrents...)
		}
	}

	return res
}

func registerScraper(scraper *scraper.Scraper) {
	scrapers[scraper.ID] = scraper
}
