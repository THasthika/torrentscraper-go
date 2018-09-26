package torrentscraper

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {

	scraper := New()

	r := scraper.Search("Elementary S01E01").FilterExcludeAny("1080", "720").FilterMatchAll("HDTV")

	for _, t := range r.Torrents {
		fmt.Println(t.Name, t.Seeds)
	}

	r = scraper.SearchMovie("Iron Man")

	for _, t := range r.Torrents {
		fmt.Println(t.Name, t.Seeds)
	}
}
