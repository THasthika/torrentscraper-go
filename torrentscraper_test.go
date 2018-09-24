package torrentscraper

import "testing"

func TestMain(t *testing.T) {
	Init()

	Search("Elementary S01E01")

	SearchMovie("Iron Man")
}
