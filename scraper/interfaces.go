package scraper

// Scraper Implementation
type Scraper struct {
	ID             string
	SupportedTypes TorrentType
	Search         func(query string, t TorrentType, out chan Result)
	SearchShow     func(name string, season uint, episode uint, out chan Result)
	SearchMovie    func(name string, out chan Result)
}

// Result result from the scraper
type Result struct {
	ID       string
	Torrents []*TorrentMeta
	Err      error
}

// TorrentType of the Torrent
type TorrentType int

// Set of TorrentTypes
const (
	TorrentTypeUnspecified = 0x00
	TorrentTypeMovie       = 0x01
	TorrentTypeTV          = 0x02
	TorrentTypeGame        = 0x04
	TorrentTypeBook        = 0x05
)

// TorrentMeta stores the information about the torrent
type TorrentMeta struct {
}
