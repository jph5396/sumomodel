package sumomodel

import (
	"fmt"
	"testing"
)

func TestCreateRankModelFromString(t *testing.T) {
	tests := []struct {
		data string
		want string
	}{
		{"S1e", "Rank: S Num: 1 Side: e Overhang: "},
		{"Y2wYO", "Rank: Y Num: 2 Side: w Overhang: YO"},
		{"Sd100w", "Rank: Sd Num: 100 Side: w Overhang: "},
		{"Ms10TD", "Rank: Ms Num: 10 Side:  Overhang: TD"},
	}

	for _, test := range tests {
		got, err := CreateRankModelFromString(test.data)
		if err != nil {
			t.Errorf(err.Error())
		}

		gotstr := fmt.Sprintf("Rank: %v Num: %v Side: %v Overhang: %v", got.Name, got.Num, got.Side, got.Overhang)
		if test.want != gotstr {
			t.Errorf("Rank did not parse correctly. \n got: %v \n want: %v", gotstr, test.want)
		}

	}
}

func TestSortRikishiByRank(t *testing.T) {
	r1 := Rikishi{Rank: "S1e"}
	r2 := Rikishi{Rank: "Ms10w"}
	r3 := Rikishi{Rank: "Ms10TD"}
	r4 := Rikishi{Rank: "J12e"}
	r5 := Rikishi{Rank: "S1w"}

	tests := []struct {
		data []Rikishi
		want []Rikishi
	}{
		{[]Rikishi{r3, r4, r1}, []Rikishi{r1, r4, r3}},
		//fail on purpose
		{[]Rikishi{r5, r1}, []Rikishi{r1, r5}},
		{[]Rikishi{r3, r2}, []Rikishi{r2, r3}},
	}

	for _, test := range tests {
		testCopy := make([]Rikishi, len(test.data))
		copy(testCopy, test.data)
		got, err := SortRikishiByRank(testCopy)
		if err != nil {
			t.Error(err.Error())
		}

		if len(test.data) != len(got) {
			t.Errorf("unequal length of arrays. \n got %v \n want: %v", len(got), len(test.data))
			break
		}

		for i, val := range got {
			if val.Rank != test.want[i].Rank {
				t.Errorf("sort order incorrect. \n got: %v \n want: %v", val.Rank, test.want[i].Rank)
				break
			}
		}
	}
}
