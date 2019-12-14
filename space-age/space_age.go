/*
* Package space provide age conversion method for different planets
*/
package space

import (
	"math"
)

// Create custom type since that what's the function gets. 
type Planet string	

// Age convert seconds to years age on diffrents planets.
func Age(ageInSeconds float64,  planet Planet) float64 {

	var plantsOrbitalPeriod = map [Planet] float64 {
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1.0,
		"Mars":    1.8808158,
		"Jupiter": 11.862615, 
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132, // Comma is necessary
	}

	 var secondsInEarthOrbit float64 = 31557600
	 var years float64 = ageInSeconds / (secondsInEarthOrbit * plantsOrbitalPeriod[planet])
	 // Round the number to match the test_case expected value.
	 var yearsRoundedNumber = math.Round(years * 100) / 100

	 // return var and not directly to make it more clear what's returning here.
	return yearsRoundedNumber
}
