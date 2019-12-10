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

	// var person = Person{
	// 	Name: "shmuel",
	// 	Address: "test",
	// }

	// var st = "Name"

	// value, err := reflections.GetField(person, st)
	//value = "test" //:= reflect.Indirect(person).FieldByName(st)

	var ageInSeconds float64 = 189839836 //1000000000
	var VenusOrbitalPeriod float64 = 0.61519726
	var secondsInEarthOrbit float64 = 31557600
	var plantOrbitalPeriod = VenusOrbitalPeriod
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