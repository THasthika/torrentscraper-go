package torrentscraper

import (
	"strings"

	"github.com/tharindu96/torrentscraper-go/scraper"
)

// FilterResultMatchAll filters torrents from the result
func FilterResultMatchAll(torrents []*scraper.TorrentMeta, keywords ...string) []*scraper.TorrentMeta {
	res := make([]*scraper.TorrentMeta, 0)

	for _, t := range torrents {
		found := true
		for _, k := range keywords {
			if strings.Contains(t.Name, k) != true {
				found = false
				break
			}
		}
		if found {
			res = append(res, t)
		}
	}

	return res
}

// FilterResultMatchAny filters torrents from the result
func FilterResultMatchAny(torrents []*scraper.TorrentMeta, keywords ...string) []*scraper.TorrentMeta {
	res := make([]*scraper.TorrentMeta, 0)

	for _, t := range torrents {
		for _, k := range keywords {
			if strings.Contains(t.Name, k) {
				res = append(res, t)
				break
			}
		}
	}

	return res
}

// FilterResultExcludeAll filters torrents from the result
func FilterResultExcludeAll(torrents []*scraper.TorrentMeta, keywords ...string) []*scraper.TorrentMeta {
	res := make([]*scraper.TorrentMeta, 0)

	for _, t := range torrents {
		found := 0
		for _, k := range keywords {
			if strings.Contains(t.Name, k) {
				found++
			}
		}
		if found != len(keywords) {
			res = append(res, t)
		}
	}

	return res
}

// FilterResultExcludeAny filters torrents from the result
func FilterResultExcludeAny(torrents []*scraper.TorrentMeta, keywords ...string) []*scraper.TorrentMeta {
	res := make([]*scraper.TorrentMeta, 0)

	for _, t := range torrents {
		found := true
		for _, k := range keywords {
			if strings.Contains(t.Name, k) {
				found = false
				break
			}
		}
		if found {
			res = append(res, t)
		}
	}

	return res
}
