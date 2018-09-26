package torrentscraper

import (
	"sort"

	"github.com/tharindu96/torrentscraper-go/scraper"
	"github.com/tharindu96/torrentscraper-go/scraper/eztvag"
	"github.com/tharindu96/torrentscraper-go/scraper/zooqlecom"
)

var scrapers map[string]*scraper.Scraper

// Init func
func Init(excludeScrapers ...string) {
	scrapers = make(map[string]*scraper.Scraper)

	registerScraper(eztvag.Init(), &excludeScrapers)
	registerScraper(zooqlecom.Init(), &excludeScrapers)
}

// Search func
func Search(query string) []*scraper.TorrentMeta {
	return searchType(query, scraper.TorrentTypeUnspecified)
}

// SearchMovie func
func SearchMovie(query string) []*scraper.TorrentMeta {
	return searchType(query, scraper.TorrentTypeMovie)
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

	tmap := make(map[string]*scraper.TorrentMeta)

	for i := 0; i < count; i++ {
		result := <-out
		if result.Err == nil && result.Torrents != nil {
			mergeTorrents(result.Torrents, &tmap)
		}
	}
	close(out)

	res := make([]*scraper.TorrentMeta, len(tmap))
	i := 0
	for _, v := range tmap {
		res[i] = v
		i++
	}

	sort.Slice(res, func(i, j int) bool {
		if res[i].Seeds < res[j].Seeds {
			return false
		}
		return true
	})

	return res
}

func mergeTorrents(torrents []*scraper.TorrentMeta, tmap *map[string]*scraper.TorrentMeta) {
	for _, t := range torrents {
		hash, err := scraper.GetTorrentXt(t.Magnet)
		if err != nil {
			continue
		}
		if x, ok := (*tmap)[hash]; ok {
			if x.Seeds < t.Seeds {
				(*tmap)[hash] = t
			}
		} else {
			(*tmap)[hash] = t
		}
	}
}

func registerScraper(scraper *scraper.Scraper, excludeScrapers *[]string) {
	found := false
	for _, s := range *excludeScrapers {
		if scraper.ID == s {
			found = true
		}
	}
	if !found {
		scrapers[scraper.ID] = scraper
	}
}
