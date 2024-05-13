package routine

import (
	"fmt"
)

var tasks = []string{"sleep", "asakatsu"}

type Routine struct {
	Name           string // 半角英数字
	Description    string
	Rank           string
	RankUpAt       string
	RankDownAt     string
	RankUpCount    string
	RankKeepCount  string
	skipticketsNum int
	UpdateAt       string
}

func ratedRoutine() {
	ranks := map[string]int{
		"SS": 3,
		"S":  5,
		"A":  3,
		"B":  3,
		"C":  5,
		"D":  3,
		"E":  3,
		"F":  1,
	}
	fmt.Println(ranks)
}
