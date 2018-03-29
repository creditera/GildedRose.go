package main

import "fmt"

// UpdateQuality accepts an collection of Items and returns a modified version
// of that collection. It implements all the rules described in the README.md
func UpdateQuality(items []Item) []Item {
	updatedItems := make([]Item, len(items))

	for idx, item := range items {
		if item.Name != "Aged Brie" && item.Name != "Backstage passes to a TAFKAL80ETC concert" {
			if item.Quality > 0 {
				if item.Name != "Sulfuras, Hand of Ragnaros" {
					item.Quality -= 1
				}
			}
		} else {
			if item.Quality < 50 {
				item.Quality += 1
				if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
					if item.SellIn < 11 {
						if item.Quality < 50 {
							item.Quality += 1
						}
					}
					if item.SellIn < 6 {
						if item.Quality < 50 {
							item.Quality += 1
						}
					}
				}
			}
		}
		if item.Name != "Sulfuras, Hand of Ragnaros" {
			item.SellIn -= 1
		}
		if item.SellIn < 0 {
			if item.Name != "Aged Brie" {
				if item.Name != "Backstage passes to a TAFKAL80ETC concert" {
					if item.Quality > 0 {
						if item.Name != "Sulfuras, Hand of Ragnaros" {
							item.Quality -= 1
						}
					}
				} else {
					item.Quality = item.Quality - item.Quality
				}
			} else {
				if item.Quality < 50 {
					item.Quality += 1
				}
			}
		}

		updatedItems[idx] = item
	}

	return updatedItems
}

/*
  WARNING: DO NOT MODIFY ANYTHING BEYOND THIS POINT

  The functionality of the remainder of this file is assumed by all other parts
  of the Gilded Rose kata.
 */

// Item describes an item sold by the Gilded Rose Inn
type Item struct {
	Name    string
	SellIn  int
	Quality int
}

func main() {
	fmt.Sprint("This program does nothing directly. Run its test suite instead.")
}