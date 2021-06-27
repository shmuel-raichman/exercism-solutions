//B"H
/*
Packege raindrops help converting raindrops as number to raindrop sounds as string.
*/
package raindrops

import (
	"strconv"
)

// Convert converts a raindrop number to pre determined sounds
func Convert(raindrops int) string {

	factorsAndSounds := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}

	// factors slice used to keep map iteration in order (3, 5, 7)
	factors := [3]int{3, 5, 7}

	var raindropsSound string

	for _, factor := range factors {
		if raindrops%factor == 0 {
			raindropsSound += factorsAndSounds[factor]
		}
	}

	// If no factors set raindropsSound string to raindrops int
	if len(raindropsSound) == 0 {
		raindropsSound = strconv.Itoa(raindrops)
	}

	return raindropsSound
}
