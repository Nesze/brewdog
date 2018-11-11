package brewery

import (
	"log"
	"strings"

	"github.com/Nesze/brewdog/model"
	"github.com/Nesze/brewdog/scraper"
)

// Brewery ...
type Brewery interface {
	Bars() []model.Bar
	OnTap(barName string) map[string][]model.Beer
	WhereIsOnTap(beerName string) []model.Bar
}

// Brewdog ...
type brewdog struct {
	bars []model.Bar
}

// Checkout ...
func Checkout() Brewery {
	var bars []model.Bar
	for k, v := range scraper.ScrapeBars() {
		bars = append(bars, model.Bar{
			Name: k,
			ID:   v,
		})
	}
	return brewdog{bars: bars}
}

// OnTap ...
func (b brewdog) OnTap(barName string) map[string][]model.Beer {
	for _, bar := range b.bars {
		if bar.Name == barName {
			return groupByLabel(scraper.ScrapeBeers(bar.ID))
		}
	}
	log.Fatal("Bar not found")
	return nil
}

func (b brewdog) Bars() []model.Bar {
	return b.bars
}

func (b brewdog) WhereIsOnTap(beerName string) []model.Bar {
	var bars []model.Bar
	for _, bar := range b.bars {
		var beers []model.Beer
		for _, beer := range scraper.ScrapeBeers(bar.ID) {
			if strings.Contains(strings.ToLower(beer.Name), beerName) {
				beers = append(beers, beer)
			}
		}
		if len(beers) != 0 {
			bar.OnTap = beers
			bars = append(bars, bar)
		}
	}
	return bars
}

func groupByLabel(beers []model.Beer) map[string][]model.Beer {
	m := make(map[string][]model.Beer)

	for _, beer := range beers {
		m[beer.Label] = append(m[beer.Label], beer)
	}
	return m
}
