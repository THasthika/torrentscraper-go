package torrentscraper

import (
	"github.com/tharindu96/torrentscraper-go/scraper"
	"github.com/tharindu96/torrentscraper-go/scraper/eztvag"
)

var scrapers map[string]*scraper.Scraper

// Init func
func Init() {
	scrapers = make(map[string]*scraper.Scraper)

	registerScraper(eztvag.Init())
}

func registerScraper(scraper *scraper.Scraper) {
	scrapers[scraper.ID] = scraper
}

func Search(query string) {

}
