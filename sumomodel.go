package sumomodel

import (
	"fmt"
)

type (
	// Rikishi data struct.
	Rikishi struct {
		Id         int
		Name       string
		Rank       string
		Heya       string
		Shusshin   string
		HW         string
		Result     string
		Kanji      string
		Dob        string
		Firstbasho string
		Lastbasho  string
		Division   int
		BashoID    int
	}

	//Bout represents a single bout between two rikishi.
	Bout struct {
		BashoID         int
		Day             int
		Boutnum         int
		Division        int
		EastRikishiID   int
		EastRikishiName string
		EastWin         bool
		WestRikishiID   int
		WestRikishiName string
		WestWin         bool
		Kimarite        string
	}
)

// PrintData prints the data for the Bout struct
func (b *Bout) PrintData() {
	fmt.Printf("day: %v Bout: %v Div: %v EastId: %v EName: %v EWin: %v WestId: %v WName: %v WWin: %v Kimarite: %v",
		b.Day,
		b.Boutnum,
		b.Division,
		b.EastRikishiID,
		b.EastRikishiName,
		b.EastWin,
		b.WestRikishiID,
		b.WestRikishiName,
		b.WestWin,
		b.Kimarite)
	fmt.Println()
}

// PrintData prints some of the rikishi structs data as a test.
func (r *Rikishi) PrintData() {
	fmt.Printf(
		"id: %v, rank: %v, name: %v, kanji: %v, heya: %v, shusshin: %v, dob = %v, results = %v",
		r.Id,
		r.Rank,
		r.Name,
		r.Kanji,
		r.Heya,
		r.Shusshin,
		r.Dob,
		r.Result)
	fmt.Println()
}
