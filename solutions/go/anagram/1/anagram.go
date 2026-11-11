package anagram

import (
	"unicode"
)

func Detect(target string, candidates []string) []string {
	var result []string

	targetRunes := []rune(target)
	var lowerTarget []rune
	targetCounts := make(map[rune]int)

	for _, r := range targetRunes {
		lr := unicode.ToLower(r)
		lowerTarget = append(lowerTarget, lr)
		targetCounts[lr]++
	}

	for _, candidate := range candidates {
		candidateRunes := []rune(candidate)

		if len(candidateRunes) != len(lowerTarget) {
			continue
		}

		var lowerCandidate []rune
		isSameWord := true
		candidateCounts := make(map[rune]int)

		for i, r := range candidateRunes {
			lr := unicode.ToLower(r)
			lowerCandidate = append(lowerCandidate, lr)
			candidateCounts[lr]++

			if lr != lowerTarget[i] {
				isSameWord = false
			}
		}

		if isSameWord {
			continue
		}

		if mapsEqual(targetCounts, candidateCounts) {
			result = append(result, candidate)
		}
	}

	return result
}

func mapsEqual(a, b map[rune]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}