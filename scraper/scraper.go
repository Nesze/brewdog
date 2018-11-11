package scraper

import (
	"log"
	"strings"

	"github.com/Nesze/brewdog/model"
	"github.com/gocolly/colly"
)

const baseURL = "https://www.brewdog.com"
const barListURL = baseURL + "/bars/uk"
const tapListURL = baseURL + "/ajax/tap_list.php?id="

// ScrapeBars ...
func ScrapeBars() map[string]string {
	bars := make(map[string]string)

	c := colly.NewCollector()
	c.OnHTML("a[href][title]", func(e *colly.HTMLElement) {
		if !strings.HasPrefix(e.Attr("href"), "/bars/uk/") {
			// ignore non-UK bars
			return
		}

		if !strings.HasPrefix(e.Attr("id"), "item") {
			// not a bar list item
			return
		}

		id := strings.TrimSuffix(strings.TrimPrefix(e.Attr("id"), "item_"), "_a")
		bars[strings.ToLower(e.Attr("title"))] = id
	})

	err := c.Visit(barListURL)
	if err != nil {
		log.Fatal(err)
	}

	return bars
}

// ScrapeBeers
func ScrapeBeers(id string) []model.Beer {
	url := tapListURL + id
	var beers []model.Beer

	c := colly.NewCollector()
	c.OnHTML("div[id=\"tap-list\"]", func(e *colly.HTMLElement) {
		e.ForEach("div[class=\"category\"]", func(i int, e *colly.HTMLElement) {
			title := e.ChildText("div[class=\"title\"]")

			e.ForEach("ul[class=\"beer\"]", func(i int, e *colly.HTMLElement) {

				var name, style, brewery, abv string
				e.ForEach("li", func(i int, e *colly.HTMLElement) {
					switch i {
					case 0:
						name = e.Text
					case 1:
						style = e.Text
					case 2:
						e.ForEach("span", func(i int, e *colly.HTMLElement) {
							switch i {
							case 0:
								brewery = e.Text
							case 1:
								abv = e.Text
							}
						})
					}
				})

				b := model.Beer{
					Name:    name,
					Style:   style,
					Brewery: brewery,
					ABV:     abv,
					Label:   title,
				}
				beers = append(beers, b)
			})
		})
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	return beers
}
