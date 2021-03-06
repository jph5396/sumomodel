package sumomodel

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type (
	// Rikishi data struct.
	// the Name variable refers to the shikona, while the RealName variable is the
	// rikishis given legal name.
	Rikishi struct {
		Id         int
		Name       string
		RealName   string
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

	// Rank is a model used to make it easier to compare rikishi.
	// it is not used in the rikishi model because strings suffice in
	// most situations.
	Rank struct {
		Name     string
		Num      int
		Side     string
		Overhang string
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

//SortRikishiByRank takes a list of Rikishi's and sorts it by rank.
func SortRikishiByRank(r []Rikishi) error {

	sort.SliceStable(r,
		func(i int, j int) bool {
			ranki, err := CreateRankModelFromString(r[i].Rank)
			if err != nil {
				panic(err)
			}
			rankj, err := CreateRankModelFromString(r[j].Rank)
			if err != nil {
				panic(err)
			}
			rankValue := map[string]int{
				"Y":  1,
				"O":  2,
				"S":  3,
				"K":  4,
				"M":  5,
				"J":  6,
				"Ms": 7,
				"Sd": 8,
				"Jd": 9,
				"Jk": 10,
			}

			sideValue := map[string]int{
				"e": 1,
				"w": 2,
				"":  3,
			}

			if rankValue[ranki.Name] != rankValue[rankj.Name] {
				return rankValue[ranki.Name] < rankValue[rankj.Name]
			} else if ranki.Num != rankj.Num {
				return ranki.Num < rankj.Num
			}
			return sideValue[ranki.Side] < sideValue[rankj.Side]

		})

	return nil

}

//SortBouts sorts bouts by bashoId, day, then bout number.
func SortBouts(b []Bout) error {

	sort.SliceStable(b,
		func(i int, j int) bool {
			bouti := b[i]
			boutj := b[j]

			if bouti.BashoID != boutj.BashoID {
				return bouti.BashoID < boutj.BashoID
			} else if bouti.Day != boutj.Day {
				return bouti.Day < boutj.Day
			}
			return bouti.Boutnum < boutj.Boutnum
		})

	return nil
}

//CreateRankModelFromString takes a string and attempts to convert
// it into a rank struct. will return an error if anything in the process
// goes wrong.
func CreateRankModelFromString(str string) (Rank, error) {
	// check if the str can be converted to a rank by calling
	// GetDivisionByRank
	_, err := GetDivisionByRank(str)
	if err != nil {
		return Rank{}, fmt.Errorf("%v can not be cast to a rank", str)
	}

	var rank Rank
	strAsRune := []rune(str)

	if unicode.IsNumber(strAsRune[1]) {
		rank.Name = string(str[0])
	} else {
		rank.Name = string([]byte{str[0], str[1]})
	}

	// use regex to get number
	re := regexp.MustCompile("[0-9]+")
	rank.Num, err = strconv.Atoi(re.FindString(str))
	if err != nil {
		return Rank{}, err
	}

	// remove the characters that we have already used from the string to
	// make it easier to parse the rest.
	str = strings.Replace(str, rank.Name, "", 1)
	str = strings.ReplaceAll(str, strconv.Itoa(int(rank.Num)), "")

	// e or w represents the east and west sides is present in all but a few
	// situations.
	if string(str[0]) == "e" || string(str[0]) == "w" {
		rank.Side = string(str[0])
		// if there are any characters left after setting rank side,
		// apply the remaining to the Overhang section.
		if len(str) > 1 {
			rank.Overhang = string(str[1:])
		}
	} else {
		rank.Overhang = str
	}

	return rank, nil
}
