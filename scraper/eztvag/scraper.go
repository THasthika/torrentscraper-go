package eztvag

import (
	"errors"

	"github.com/tharindu96/torrentscraper-go/scraper"
)

const id = "eztvag"
const ttype = scraper.TorrentTypeTV

// Init func
func Init() *scraper.Scraper {
	return &scraper.Scraper{
		ID:             id,
		SupportedTypes: ttype,
		Search:         Search,
		SearchShow:     SearchShow,
	}
}

// Search func
func Search(query string, t scraper.TorrentType, out chan scraper.Result) {
	if t&ttype == 0 {
		out <- scraper.Result{
			ID:       id,
			Torrents: nil,
			Err:      errors.New("asf"),
		}
		return
	}

}

// SearchShow func
func SearchShow(name string, season uint, episode uint, out chan scraper.Result) {

}
