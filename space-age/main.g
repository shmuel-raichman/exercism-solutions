// B"H
package main
import (
	"fmt"
	"math"
	// "reflect"
)

// {
// 	description: "age on Earth",
// 	planet:      "Earth",
// 	seconds:     1000000000,
// 	expected:    31.69,
// },
// {
// 	description: "age on Mercury",
// 	planet:      "Mercury",
// 	seconds:     2134835688,
// 	expected:    280.88,
// },
// {
// 	description: "age on Venus",
// 	planet:      "Venus",
// 	seconds:     189839836,
// 	expected:    9.78,

// func plantOrbitalPeriod(planet) float64 {
	
// 	mercury = 0.2408467
// 	venus = 0.61519726
// 	earth = 1.0
// 	mars = 1.8808158
// 	jupiter = 11.862615
// 	saturn = 29.447498
// 	uranus = 84.016846
// 	neptune = 164.79132s

// 	return plantOrbitalPeriod;
// }


// type Person struct {
// 	Name string
// 	Address string
//  }


func main(){

	// type PlantPeriod float64

	// const (
	// 	Mercury PlantPeriod = 0.2408467
	// 	Venus   PlantPeriod = 0.61519726
	// 	Earth   PlantPeriod = 1.0
	// 	Mars    PlantPeriod = 1.8808158
	// 	Jupiter PlantPeriod = 11.862615
	// 	Saturn  PlantPeriod = 29.447498
	// 	Uranus  PlantPeriod = 84.016846
	// 	Neptune PlantPeriod = 164.79132
	//  )

	var planets = map[string]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1.0,
		"Mars":    1.8808158,
		"Jupiter": 11.862615, 
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132, // Comma is necessary'
	}

	//  - Mercury: orbital period 0.2408467 Earth years
	//  - Venus: orbital period 0.61519726 Earth years
	//  - Earth: orbital period 1.0 Earth years, 365.25 Earth days, or 31557600 seconds
	//  - Mars: orbital period 1.8808158 Earth years
	//  - Jupiter: orbital period 11.862615 Earth years
	//  - Saturn: orbital period 29.447498 Earth years
	//  - Uranus: orbital period 84.016846 Earth years
	//  - Neptune: orbital period 164.79132 Earth years
	// var person = Person{
	// 	Name: "shmuel",
	// 	Address: "test",
	// }

	var planet = "Earth"

	// value, err := reflections.GetField(person, st)
	//value = "test" //:= reflect.Indirect(person).FieldByName(st)
	// Mercury = 2134835688 = 280.88
	// Venus = 189839836 = 9.78
	// Earth = 1000000000 = 31.69
	var ageInSeconds float64 = 1000000000 //1000000000
	// var VenusOrbitalPeriod float64 = 0.61519726
	var secondsInEarthOrbit float64 = 31557600
	var plantOrbitalPeriod float64 = planets[planet]
	// var planet = "fasd"
	// // 60 x 60 x 24 x 365
	// var age = 1.1
	// var hour float64 = 60
	// var second float64 = 60
	// var day float64 = 24
	// var year float64 = 365.25
	// var VenusOrbitalPeriod float64 = 0.61519726
	var years float64 = ageInSeconds / (secondsInEarthOrbit * plantOrbitalPeriod)
	// var secondsInEarth = seconds / (31557600 * VenusOrbitalPeriod)
	fmt.Printf("Debug: %#v \n", years)
	//var years float64 = earthSeconds / (year * day * hour * second)
	var yearsRoundedNumber = math.Round(years * 100) / 100

	fmt.Printf("Debug: %#v ", yearsRoundedNumber)
}