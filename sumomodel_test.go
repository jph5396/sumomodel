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
