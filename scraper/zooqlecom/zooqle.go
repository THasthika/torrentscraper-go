package zooqlecom

import (
	"fmt"
	"strconv"

	"github.com/PuerkitoBio/goquery"

	"github.com/tharindu96/torrentscraper-go/scraper"
)

const id = "zooqlecom"
const ttype = scraper.TorrentTypeTV | scraper.TorrentTypeBook | scraper.TorrentTypeGame | scraper.TorrentTypeMovie

const urlPlaceholder = "https://zooqle.com/search?q=%s"

const colName = 1
const colLink = 2
const colSize = 3
const colSeeds = 5

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

	res := scraper.Result{
		ID: id,
	}

	if t&ttype != t {
		out <- res
		return
	}

	ret, err := search(query)

	res.Torrents = ret
	res.Err = err

	out <- res
}

// SearchShow func
func SearchShow(name string, season uint, episode uint, out chan scraper.Result) {
	var query string
	if episode > 0 {
		query = fmt.Sprintf("%s-s%02de%02d", name, season, episode)
	} else {
		query = fmt.Sprintf("%s-s%02d", name, season)
	}

	ret, err := search(query)

	res := scraper.Result{
		ID:       id,
		Torrents: ret,
		Err:      err,
	}

	out <- res
}

func search(query string) (torrents []*scraper.TorrentMeta, err error) {
	ret := make([]*scraper.TorrentMeta, 0)

	url := fmt.Sprintf(urlPlaceholder, query)

	doc, err := scraper.GetGoQueryDocument(url)
	if err != nil {
		return nil, err
	}

	table := doc.Find("table.table-torrents tbody")
	table.Find("tr").Each(func(i int, row *goquery.Selection) {
		t := scraper.TorrentMeta{}
		row.Find("td").Each(func(j int, col *goquery.Selection) {
			switch j {
			case colName:
				t.Name = col.Find("a").Text()
				break
			case colLink:
				t.Magnet, _ = col.Find("a[title=\"Magnet link\"]").Attr("href")
				break
			case colSeeds:
				s, err := strconv.ParseUint(col.Find("div.prog-green").Text(), 10, 32)
				if err == nil {
					t.Seeds = uint(s)
				}
				break
			case colSize:
				size, err := scraper.ConvertSize(col.Text())
				if err == nil {
					t.Size = size
				}
				break
			}
		})
		ret = append(ret, &t)
	})

	return ret, nil
}
