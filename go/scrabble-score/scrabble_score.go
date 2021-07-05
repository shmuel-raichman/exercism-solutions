// B"H
/*
Package scrabble provide scoring functionality for scrabble game.
*/
package scrabble

import "strings"

// Score give a score to given word acording to a map of scores.
func AScore(word string) int {
	scoreMap := map[string]int{
		"A": 1, "E": 1, "I": 1, "O": 1, "U": 1, "L": 1, "N": 1, "R": 1, "S": 1, "T": 1,
		"D": 2, "G": 2,
		"B": 3, "C": 3, "M": 3, "P": 3,
		"F": 4, "H": 4, "V": 4, "W": 4, "Y": 4,
		"K": 5,
		"J": 8, "X": 8,
		"Q": 10, "Z": 10,
	}

	var score int

	for pos := range word {
		// Get the current char as upper case by
		// converting the current rune to string and then to upper case.
		currentChar := strings.ToUpper(string(word[pos]))

		// Add current char score to total word score.
		score += scoreMap[currentChar]

	}

	return score
}

// Score give a score to given word acording to a map of scores.
func BScore(word string) int {

	var score int

	for _, char := range word {

		switch char {
		case 'A', 'a', 'E', 'e', 'I', 'i', 'O', 'o', 'U', 'u', 'L', 'l', 'N', 'n', 'R', 'r', 'S', 's', 'T', 't':
			score += 1
		case 'D', 'd', 'G', 'g':
			score += 2
		case 'B', 'b', 'C', 'c', 'M', 'm', 'P', 'p':
			score += 3
		case 'F', 'f', 'H', 'h', 'V', 'v', 'W', 'w', 'Y', 'y':
			score += 4
		case 'K', 'k':
			score += 5
		case 'J', 'j', 'X', 'x':
			score += 8
		case 'Q', 'q', 'Z', 'z':
			score += 10
		default:
			println("Non")
		}

	}

	return score
}

// Score give a score to given word acording to a map of scores.
func CScore(word string) int {
	scoreMap := map[rune]int{
		'A': 1, 'a': 1, 'E': 1, 'e': 1, 'I': 1, 'i': 1, 'O': 1, 'o': 1, 'U': 1, 'u': 1, 'L': 1, 'l': 1, 'N': 1, 'n': 1, 'R': 1, 'r': 1, 'S': 1, 's': 1, 'T': 1, 't': 1,
		'D': 2, 'd': 2, 'G': 2, 'g': 2,
		'B': 3, 'b': 3, 'C': 3, 'c': 3, 'M': 3, 'm': 3, 'P': 3, 'p': 3,
		'F': 4, 'f': 4, 'H': 4, 'h': 4, 'V': 4, 'v': 4, 'W': 4, 'w': 4, 'Y': 4, 'y': 4,
		'K': 5, 'k': 5,
		'J': 8, 'j': 8, 'X': 8, 'x': 8,
		'Q': 10, 'q': 10, 'Z': 10, 'z': 10,
	}

	var score int

	for _, char := range word {
		// Get the current char as upper case by
		// converting the current rune to string and then to upper case.
		// currentChar := strings.ToUpper(string(word[pos]))

		// Add current char score to total word score.
		score += scoreMap[char]

	}

	return score
}
