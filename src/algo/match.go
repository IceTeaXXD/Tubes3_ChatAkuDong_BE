package Algo

import "github.com/texttheater/golang-levenshtein/levenshtein"

func MatchRatio(a string, b string) float64 {
	return levenshtein.RatioForStrings([]rune(a), []rune(b), levenshtein.DefaultOptions)
}