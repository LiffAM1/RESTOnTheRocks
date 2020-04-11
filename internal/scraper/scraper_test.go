package scraper

import (
	"testing"
)

func TestScraper(t *testing.T) {
	var recipes = ScrapeByName("anything")
	for i := 0; i < len(recipes); i++ {
		t.Log(recipes[i].Name)
	}
}