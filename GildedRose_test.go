package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/plainprogrammer/GildedRose"
)

func CreateAndUpdateItem(name string, sellIn int, quality int) Item {
	item := Item{Name: name, SellIn: sellIn, Quality: quality}
	items := UpdateQuality([]Item{item})
	return items[0]
}

var _ = Describe("GildedRose", func() {
	Describe("UpdateQuality", func() {
		Context("with a single", func() {
			Context("normal item", func() {
				productName := "Normal Item"

				It("decrements sell in by 1", func() {
					item := CreateAndUpdateItem(productName, 5, 10)
					Expect(item.SellIn).To(Equal(4))
				})

				Context("before the sell date", func() {
					item := CreateAndUpdateItem(productName, 5, 10)

					It("degrades quality by 1", func() {
						Expect(item.Quality).To(Equal(9))
					})
				})

				Context("on the sell date", func() {
					item := CreateAndUpdateItem(productName, 0, 10)

					It("degrades quality by 2", func() {
						Expect(item.Quality).To(Equal(8))
					})
				})

				Context("after sell date", func() {
					item := CreateAndUpdateItem(productName, -10, 10)

					It("degrades quality by 2", func() {
						Expect(item.Quality).To(Equal(8))
					})
				})

				Context("of zero quality", func() {
					item := CreateAndUpdateItem(productName, 5, 0)

					It("does not decrement the quality below 0", func() {
						Expect(item.Quality).To(Equal(0))
					})
				})
			})

			Context("Aged Brie", func() {
				productName := "Aged Brie"

				It("decrements sell in by 1", func() {
					item := CreateAndUpdateItem(productName, 5, 10)
					Expect(item.SellIn).To(Equal(4))
				})

				Context("before the sell date", func() {
					item := CreateAndUpdateItem(productName, 5, 10)

					It("increases quality by 1", func() {
						Expect(item.Quality).To(Equal(11))
					})

					Context("with max quality", func() {
						item := CreateAndUpdateItem(productName, 5, 50)

						It("does not increase quality beyond 50", func() {
							Expect(item.Quality).To(Equal(50))
						})
					})
				})

				Context("on the sell date", func() {
					item := CreateAndUpdateItem(productName, 0, 10)

					It("increases quality by 2", func() {
						Expect(item.Quality).To(Equal(12))
					})

					Context("near max quality", func() {
						item := CreateAndUpdateItem(productName, 0, 49)

						It("does not increase quality beyond 50", func() {
							Expect(item.Quality).To(Equal(50))
						})
					})

					Context("with max quality", func() {
						item := CreateAndUpdateItem(productName, 0, 50)

						It("does not increase quality beyond 50", func() {
							Expect(item.Quality).To(Equal(50))
						})
					})
				})

				Context("after sell date", func() {
					item := CreateAndUpdateItem(productName, -10, 10)

					It("increases quality by 2", func() {
						Expect(item.Quality).To(Equal(12))
					})

					Context("with max quality", func() {
						item := CreateAndUpdateItem(productName, -10, 50)

						It("does not increase quality beyond 50", func() {
							Expect(item.Quality).To(Equal(50))
						})
					})
				})
			})

			Context("Sulfuras, Hand of Ragnaros", func() {
				productName := "Sulfuras, Hand of Ragnaros"

				It("sell in does not change", func() {
					item := CreateAndUpdateItem(productName, 5, 80)
					Expect(item.SellIn).To(Equal(5))
				})

				Context("before the sell date", func() {
					item := CreateAndUpdateItem(productName, 5, 80)

					It("quality does not change", func() {
						Expect(item.Quality).To(Equal(80))
					})
				})

				Context("on the sell date", func() {
					item := CreateAndUpdateItem(productName, 0, 80)

					It("quality does not change", func() {
						Expect(item.Quality).To(Equal(80))
					})
				})

				Context("after sell date", func() {
					item := CreateAndUpdateItem(productName, -10, 80)

					It("quality does not change", func() {
						Expect(item.Quality).To(Equal(80))
					})
				})
			})

			Context("Backstage passes to a TAFKAL80ETC concert", func() {
				productName := "Backstage passes to a TAFKAL80ETC concert"

				It("decrements sell in by 1", func() {
					item := CreateAndUpdateItem(productName, 5, 10)
					Expect(item.SellIn).To(Equal(4))
				})

				Context("long before the sell date", func() {
					item := CreateAndUpdateItem(productName, 11, 10)

					It("increases quality by 1", func() {
						Expect(item.Quality).To(Equal(11))
					})

					Context("with max quality", func() {
						item := CreateAndUpdateItem(productName, 11, 50)

						It("does not increase quality beyond 50", func() {
							Expect(item.Quality).To(Equal(50))
						})
					})
				})

				Context("moderately close to the sell date (10 days - upper bound)", func() {
					item := CreateAndUpdateItem(productName, 10, 10)

					It("increases quality by 2", func() {
						Expect(item.Quality).To(Equal(12))
					})

					Context("with max quality", func() {
						item := CreateAndUpdateItem(productName, 10, 50)

						It("does not increase quality beyond 50", func() {
							Expect(item.Quality).To(Equal(50))
						})
					})
				})

				Context("moderately close to the sell date (6 days - lower bound)", func() {
					item := CreateAndUpdateItem(productName, 6, 10)

					It("increases quality by 2", func() {
						Expect(item.Quality).To(Equal(12))
					})

					Context("with max quality", func() {
						item := CreateAndUpdateItem(productName, 6, 50)

						It("does not increase quality beyond 50", func() {
							Expect(item.Quality).To(Equal(50))
						})
					})
				})

				Context("very close to the sell date (5 days - upper bound)", func() {
					item := CreateAndUpdateItem(productName, 5, 10)

					It("increases quality by 3", func() {
						Expect(item.Quality).To(Equal(13))
					})

					Context("with max quality", func() {
						item := CreateAndUpdateItem(productName, 5, 50)

						It("does not increase quality beyond 50", func() {
							Expect(item.Quality).To(Equal(50))
						})
					})
				})

				Context("very close to the sell date (5 days - lower bound)", func() {
					item := CreateAndUpdateItem(productName, 1, 10)

					It("increases quality by 3", func() {
						Expect(item.Quality).To(Equal(13))
					})

					Context("with max quality", func() {
						item := CreateAndUpdateItem(productName, 1, 50)

						It("does not increase quality beyond 50", func() {
							Expect(item.Quality).To(Equal(50))
						})
					})
				})

				Context("on the sell date", func() {
					item := CreateAndUpdateItem(productName, 0, 10)

					It("quality drops to 0", func() {
						Expect(item.Quality).To(Equal(0))
					})
				})

				Context("after sell date", func() {
					item := CreateAndUpdateItem(productName, -10, 10)

					It("quality drops to 0", func() {
						Expect(item.Quality).To(Equal(0))
					})
				})
			})

			XContext("conjured items", func() {
				productName := "Conjured Mana Cake"

				It("decrements sell in by 1", func() {
					item := CreateAndUpdateItem(productName, 5, 10)
					Expect(item.SellIn).To(Equal(4))
				})

				Context("before the sell date", func() {
					item := CreateAndUpdateItem(productName, 5, 10)

					It("degrades quality by 1", func() {
						Expect(item.Quality).To(Equal(8))
					})

					Context("of zero quality", func() {
						item := CreateAndUpdateItem(productName, 5, 0)

						It("does not decrement the quality below 0", func() {
							Expect(item.Quality).To(Equal(0))
						})
					})
				})

				Context("on the sell date", func() {
					item := CreateAndUpdateItem(productName, 0, 10)

					It("degrades quality by 2", func() {
						Expect(item.Quality).To(Equal(6))
					})

					Context("of zero quality", func() {
						item := CreateAndUpdateItem(productName, 0, 0)

						It("does not decrement the quality below 0", func() {
							Expect(item.Quality).To(Equal(0))
						})
					})
				})

				Context("after sell date", func() {
					item := CreateAndUpdateItem(productName, -10, 10)

					It("degrades quality by 2", func() {
						Expect(item.Quality).To(Equal(6))
					})

					Context("of zero quality", func() {
						item := CreateAndUpdateItem(productName, -10, 0)

						It("does not decrement the quality below 0", func() {
							Expect(item.Quality).To(Equal(0))
						})
					})
				})
			})
		})

		Context("with several items", func() {
			It("adjusts item sell in and quality appropriately", func() {
				items := []Item{
					{Name: "Normal Item", SellIn: 5, Quality: 10},
					{Name: "Aged Brie", SellIn: 3, Quality: 10},
				}
				items = UpdateQuality(items)

				Expect(items[0].Quality).To(Equal(9))
				Expect(items[0].SellIn).To(Equal(4))
				Expect(items[1].Quality).To(Equal(11))
				Expect(items[1].SellIn).To(Equal(2))
			})
		})
	})
})
