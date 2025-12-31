package main

import (
	"fmt"
	"temperature_conversions/tempconv"
)

func main() {
	//initialize a temperature in celsius
	var c tempconv.Celsius = 25.0
	//declare a variable of types Fahrenheit and Kelvin
	var f tempconv.Fahrenheit
	var k tempconv.Kelvin
	//convert the Celsius temperature to Fahrenheit and Kelvin
	f = tempconv.CToF(c)
	k = tempconv.CToK(c)
	//print the temperatures
	fmt.Println(c)
	fmt.Println(f)
	fmt.Println(k)
}
