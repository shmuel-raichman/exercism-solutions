package space

func Age(seconds, planet) float64 {

	// var age = 1.1

	type PlantPeriod float64

	const (
		Mercury PlantPeriod = 0.2408467
		Venus   PlantPeriod = 0.61519726
		Earth   PlantPeriod = 1.0
		Mars    PlantPeriod = 1.8808158
		Jupiter PlantPeriod = 11.862615
		Saturn  PlantPeriod = 29.447498
		Uranus  PlantPeriod = 84.016846
		Neptune PlantPeriod = 164.79132
	 )


	 var ageInSeconds float64 = 189839836 //1000000000
	 // var VenusOrbitalPeriod float64 = 0.61519726
	 var secondsInEarthOrbit float64 = 31557600
	 var plantOrbitalPeriod = planet

	 var years float64 = ageInSeconds / (secondsInEarthOrbit * plantOrbitalPeriod)
	 // var secondsInEarth = seconds / (31557600 * VenusOrbitalPeriod)
	 fmt.Printf("Debug: %#v \n", years)
	 //var years float64 = earthSeconds / (year * day * hour * second)
	 var yearsRoundedNumber = math.Round(years * 100) / 100
 
	 fmt.Printf("Debug: %#v ", yearsRoundedNumber)

	return age
}