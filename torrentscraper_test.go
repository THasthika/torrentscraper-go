package torrentscraper

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	Init()

	fmt.Println(Search("Elementary S01E01"))

	fmt.Println(SearchMovie("Iron Man"))
}
