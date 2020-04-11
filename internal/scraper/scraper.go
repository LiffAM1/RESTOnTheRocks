package scraper

import (
	"fmt"
	"github.com/LiffAM1/rotr/internal/dtos"
	"github.com/gocolly/colly"
	"strconv"
	"strings"
)

func ScrapeByQuery(query string, count int) []dtos.Recipe {
	queryFormatted := strings.TrimSpace(strings.Join(strings.Split(query," "),"+"))
	url := "https://www.google.com/search?q=" + queryFormatted + "&as_sitesearch=allrecipes.com/recipe"
	return scrapeGoogle(url,count)
}

// TODO: Ingredient search
// func ScrapeByIngredient(ingrName string) []dtos.Recipe

func scrapeGoogle(url string, count int) []dtos.Recipe {
	recipes := []dtos.Recipe{}
	c := colly.NewCollector(colly.AllowedDomains("www.google.com"), colly.MaxDepth(5))

	// Add site-specific collectors here
	allRecipesCollector := colly.NewCollector() // allrecipes.com

	// Google search scraper
	r := dtos.Recipe{}
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if len(recipes) >= count { return }
		foundURL := e.Request.AbsoluteURL(e.Attr("href"))
		if strings.HasPrefix(foundURL, "https://www.google.com/url?q=https://www.allrecipes.com/recipe") &&
			!strings.Contains(foundURL,"print") && !strings.Contains(foundURL,"reviews") {
			r.Url = strings.TrimPrefix(strings.Split(foundURL,"&sa")[0],"https://www.google.com/url?q=")
			allRecipesCollector.Visit(r.Url)
		}
	})


	// AllRecipes collector behavior
	allRecipesCollector.OnHTML(`#main-content`, func(e *colly.HTMLElement) {
		r.Name = e.ChildText("#recipe-main-content")
		e.ForEach("span[itemprop=recipeIngredient]", func(_ int, s *colly.HTMLElement) {
			// TODO Parse/convert ingredients to ounces, add to recipe, add ingredients
		})
		recipes = append(recipes,r)
	})

	for i := 0; len(recipes) < count; i++ {
		fmt.Println("Searching Google Page# " + strconv.Itoa(i+1))
		j := i*10
		startParam := "&start=" + strconv.Itoa(j)
		c.Visit(url + startParam)
		c.Wait()
	}
	return recipes
}

