package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type Word struct {
	word  string
	count int
}

// var re = regexp.MustCompile("[a-z0-9а-я-ÄÖÜäöüß]+")

func Top10(str string) []string {
	count := 1
	s := strings.Fields(str)
	mWords := make(map[string]int)

	for _, w := range s {
		if w != "" {
			mWords[w] += count
		}
	}

	sWords := make([]Word, 0, len(mWords))

	for k := range mWords {
		sWords = append(sWords, Word{
			word:  k,
			count: mWords[k],
		})
	}

	sort.Slice(sWords, func(i, j int) bool {
		iv, jv := sWords[i], sWords[j]
		switch {
		case iv.count != jv.count:
			return iv.count > jv.count
		default:
			return iv.word < jv.word
		}
	})

	res := make([]string, 0)
	for i, v := range sWords {
		res = append(res, v.word)
		if i > 8 {
			break
		}
	}
	return res
}
