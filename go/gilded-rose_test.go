package main

import (
	"math"
	"testing"

	"github.com/matryer/is"
)

func Test_GildedRose(t *testing.T) {
	// BEFORE - gather details needed later
	originalQuality := make(map[int]int, 0)

	for index, item := range items {
		originalQuality[index] = item.quality
	}

	// EXECUTE
	GildedRose()

	// AFTER - assertions
	i := is.New(t)

	// From spec/README: Once the sell by date has passed, quality degrades twice as fast.
	// James' note: folded into the other assertions, where applicable.

	// From spec/README: The quality of an item is never negative.
	for _, item := range items {
		i.True(item.quality >= 0)
	}

	// From spec/README: "Aged Brie" actually increases in quality the older it gets.
	i.True(float64(items[1].quality) == math.Min(50.0, float64(originalQuality[1]+1)))

	// From spec/README: The quality of an item is never more than 50.
	for _, item := range items {
		if item.name != "Sulfuras, Hand of Ragnaros" {
			i.True(item.quality <= 50)
		}
	}

	// From spec/README: "Sulfuras", being a legendary item, never has to be sold or decreases in quality.
	i.True(items[3].quality == originalQuality[3])

	// From spec/README: "Backstage passes", like aged brie, increases in quality as its sell-in value approaches;
	// quality increases by 2 when there are 10 days or less and by 3 when there are 5 days or less but quality drops
	// to 0 after the concert.
	var expBackstage int
	switch {
	case items[4].sellIn < 0:
		expBackstage = 0
	case items[4].sellIn <= 5:
		expBackstage = originalQuality[4] + 3
	case items[4].sellIn <= 10:
		expBackstage = originalQuality[4] + 2
	default:
		expBackstage = originalQuality[4] + 1
	}
	i.True(items[4].quality == expBackstage)

	// From spec/README: "Conjured" items degrade in quality twice as fast as normal items.
	// From spec/README: Once the sell by date has passed, quality degrades twice as fast.
	if items[5].sellIn < 0 {
		i.True(items[5].quality == originalQuality[5]-4)
	} else {
		i.True(items[5].quality == originalQuality[5]-2)
	}

	// From spec/README: Just for clarification, an item can never have its quality increase above 50, however
	// "Sulfuras" is a legendary item and as such its quality is 80 and it never alters.
	i.Equal(80, items[3].quality)

	// James' note: the +Dex vest and Mongoose elixir haven't had a look in yet, so we still need to cover those off.
	// From spec/README: Once the sell by date has passed, quality degrades twice as fast.
	for index := 0; index <= 2; index += 2 {
		if items[index].sellIn < 0 {
			i.True(items[index].quality == originalQuality[index]-2)
		} else {
			i.True(items[index].quality == originalQuality[index]-1)
		}
	}
}
