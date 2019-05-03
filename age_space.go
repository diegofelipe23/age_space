package main

import (
	"fmt"
)

const secondPeriodEarth float64 = 31557600

var (
	Earth   = 1.0
	Mercury = 0.2408467
	Venus   = 0.61519726
	Mars    = 1.8808158
	Jupiter = 11.862615
	Saturn  = 29.447498
	Uranus  = 84.016846
	Neptune = 164.79132
)

func main() {
	var seconds float64
	fmt.Println("Write your age in seconds: ")
	fmt.Scanf("%f", &seconds)
	fmt.Println("Your age on Earth is: ", calculateAge(seconds, Earth), "years old")
	fmt.Println("Your age on Mercury is: ", calculateAge(seconds, Mercury), "years old")
	fmt.Println("Your age on Venus is: ", calculateAge(seconds, Venus), "years old")
	fmt.Println("Your age on Mars is: ", calculateAge(seconds, Mars), "years old")
	fmt.Println("Your age on Jupiter is: ", calculateAge(seconds, Jupiter), "years old")
	fmt.Println("Your age on Saturn is: ", calculateAge(seconds, Saturn), "years old")
	fmt.Println("Your age on Uranus is: ", calculateAge(seconds, Uranus), "years old")
	fmt.Println("Your age on Neptune is: ", calculateAge(seconds, Neptune), "years old")

}

func calculateAge(seconds float64, planet float64) float64 {
	var ageEarth float64 = (seconds / secondPeriodEarth)
	var calculation float64 = (ageEarth / planet)
	return calculation
}
