package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type Word struct {
	word  string
	count int
}

var re = regexp.MustCompile("[a-z0-9а-я-ÄÖÜäöüß]+")

func Top10(str string) []string {
	sWords := make([]Word, 0)
	s := strings.Fields(str)

	for _, w := range s {
		wl := strings.ToLower(w)

		if wl != "-" && wl != "" {
			wr := re.FindString(wl)

			if !contains(sWords, wr) {
				sWords = append(sWords, Word{
					word:  wr,
					count: 1,
				})
			}
		}
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
	for _, v := range sWords {
		res = append(res, v.word)
	}
	if len(res) > 10 {
		return res[0:10]
	}
	return res
}

func contains(s []Word, str string) bool {
	for i, v := range s {
		if v.word == str {
			s[i].count++
			return true
		}
	}
	return false
}
