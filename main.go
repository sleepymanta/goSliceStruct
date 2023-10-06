package main

import "fmt"

type book struct {
	isbn  int
	title string
	year  int
}

func main() {
	var nilSlice []int
	emptySlice1 := make([]int, 0)
	emptySlice2 := []int{}

	fmt.Println(nilSlice == nil)    // true
	fmt.Println(emptySlice1 == nil) // false
	fmt.Println(emptySlice2 == nil) // false

	harryPotterBook := []book{}

	vol1 := book{
		isbn:  9781408898154,
		title: "Harry Potter And The Chamber Of Secrets",
		year:  1998,
	}
	fmt.Println(vol1.isbn)
	// 9781408898154

	harryPotterBook = append(harryPotterBook, vol1)
	harryPotterBook = append(harryPotterBook, book{
		isbn:  9781408855911,
		title: "Harry Potter And The Prisoner Of Azkaban",
		year:  1999,
	})
	harryPotterBook = append(harryPotterBook, book{
		9781338299175,
		"Harry Potter And The Goblet Of Fire",
		2000,
	})

	var vol4 *book = new(book)
	vol4.isbn = 9781408845684
	vol4.title = "Harry Potter And The Order Of The Phoenix"
	vol4.year = 2003
	harryPotterBook = append(harryPotterBook, *vol4)

	fmt.Println(harryPotterBook)
	// [{9781408898154 Harry Potter And The Chamber Of Secrets 1998} {978
	// 1408855911 Harry Potter And The Prisoner Of Azkaban 1999} {9781338
	// 299175 Harry Potter And The Goblet Of Fire 2000} {9781408845684 Ha
	// rry Potter And The Order Of The Phoenix 2003}]

	fmt.Println(harryPotterBook[0].year)
	// 1998
}
