package main

import (
	"fmt"
	"go_basics/celciusToFahrenheitConverter"
	"go_basics/circleareacalculater"
	"go_basics/oddoreven"
)

func temperatureConverterDemo() {
	var userInput float64
	fmt.Print(`Type a tempeture in C to convert it to F:`)
	_, err := fmt.Scanln(&userInput)

	if err != nil {

		for {
			fmt.Print(`Incorrect inputed value, pleases input numbers... : `)
			fmt.Scanln()
			_, err := fmt.Scanln(&userInput)

			if err == nil {
				break
			}
		}
	}

	var convertedValue float64 = celciusToFahrenheitConverter.Converter(userInput)
	fmt.Printf(`Converted %.2fC to %.2fF`, userInput, convertedValue)
}

func circleAreaDemo() {
	var userInput float64

	fmt.Print("Input a circle radius (cm): ")

	_, err := fmt.Scanln(&userInput)

	if err != nil {

		for {

			fmt.Print("Please provide a numurical value for the radius: ")

			fmt.Scanln(&userInput)
			_, err := fmt.Scanln(&userInput)
			if err == nil {
				break
			}
		}
	}

	circleArea := circleareacalculater.CalculateArea(userInput)
	fmt.Printf(`The area for the circle with radius: %.2fcm is %.2fcm2`, userInput, circleArea)
}

func oddOrEvenDemo() {

	fmt.Print(`Input a random number (press: "x" to quite anymoment): `)

	var userInput int
	_, err := fmt.Scanln(&userInput)
	if err != nil {
		for {
			fmt.Print(`Please insert a random number, not a character: `)

			fmt.Scanln()
			_, err := fmt.Scanln(&userInput)
			if err == nil {
				break
			}
		}
	}

	result := oddoreven.OddOrEven(userInput)

	fmt.Printf(`%d is even: %t`, userInput, result)
	fmt.Println()
}

func main() {
	// temperatureConverterDemo()
	// circleAreaDemo()
	// oddOrEvenDemo()
}
