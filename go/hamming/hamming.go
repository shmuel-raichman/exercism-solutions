// B"H
/*
Package hamming provide a function to measure hamming distance in strings
*/
package hamming

import "errors"

// Distance gets two strings and return the their hamming distance or error if they aren't comparable
func Distance(a, b string) (int, error) {
	var distance int = 0

	if len(a) != len(b) {
		return 0, errors.New("given strings are not equal length")
	}

	for position := range a {
		if a[position] != b[position] {
			distance++
		}
	}

	return distance, nil
}
