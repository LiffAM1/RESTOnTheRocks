package scraper

import (
	"fmt"
	"github.com/LiffAM1/rotr/internal/dtos"
	"github.com/gocolly/colly"
	"strings"
)

func ScrapeByName(recipeName string) []dtos.Recipe {
	// Search Google for the recipe name against known sites
	// For each search, go to first 10 links
	// For each link, based on the site, choose a colly collector and scrape the site and create a recipe object
	// return list of recipe objects
	recipes := scrapeGoogle("https://www.google.com/search?q=gin+and+tonic&as_sitesearch=allrecipes.com")
	for i := 0; i < len(recipes); i++ {
		fmt.Println(recipes[i].Name)
	}
	return recipes
}

func scrapeGoogle(url string) []dtos.Recipe {
	recipes := []dtos.Recipe{}
	c := colly.NewCollector(colly.AllowedDomains("www.google.com"), colly.MaxDepth(5))

	// Add site-specific collectors here
	allRecipesCollector := colly.NewCollector() // allrecipes.com

	// Google search scraper
	r := dtos.Recipe{}
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if len(recipes) > 10 { return }
		foundURL := e.Request.AbsoluteURL(e.Attr("href"))
		if strings.HasPrefix(foundURL, "https://www.google.com/url?q=https://www.allrecipes.com/recipe") &&
			!strings.Contains(foundURL,"print") && !strings.Contains(foundURL,"reviews"){
			r.Url = strings.TrimPrefix(strings.Split(foundURL,"&sa")[0],"https://www.google.com/url?q=")
			fmt.Println(r.Url)
			allRecipesCollector.Visit(r.Url)
		}
	})


	// AllRecipes collector behavior
	allRecipesCollector.OnHTML(`#main-content`, func(e *colly.HTMLElement) {
		r.Name = e.ChildText("#recipe-main-content")
		fmt.Println(r.Name)
		e.ForEach("span[itemprop=recipeIngredient]", func(_ int, s *colly.HTMLElement) {
			// TODO Parse/convert ingredients to ounces, add to recipe,add ingredients
			fmt.Println(s.Text)
		})
		recipes = append(recipes,r)
		fmt.Println(len(recipes))
	})

	for i := 0;
	c.Visit(url)
	return recipes
}

