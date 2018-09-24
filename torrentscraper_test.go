package torrentscraper

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	Init()

	x := Search("Elementary S02E01")

	for k, v := range x {
		fmt.Println(k)
		for _, y := range v {
			fmt.Println(y.Name, y.Size, y.Seeds)
		}
	}

	// fmt.Println(SearchMovie("Iron Man"))
}
