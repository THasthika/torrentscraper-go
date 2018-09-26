package torrentscraper

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	Init()

	x := Search("Elementary S01E01")
	// x = FilterResultExcludeAny(x, "1080", "720")
	// x = FilterResultMatchAll(x, "HDTV", "S01", "E01")

	for _, t := range x {
		fmt.Println(t.Name, t.Seeds)
	}

	// fmt.Println(SearchMovie("Iron Man"))
}
