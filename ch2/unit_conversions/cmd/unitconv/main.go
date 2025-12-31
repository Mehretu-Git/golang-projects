package main

import (
	"fmt"

	"os"

	"unit_conversions/pkg/conv"
)

func main() {
	//welcome message
	fmt.Println("Welcome to the Unit Conversion Tool!")
	//prompt the user to choose a conversion type
	for {
		fmt.Println("Choose a conversion type:")
		fmt.Println("1: Celsius to Fahrenheit")
		fmt.Println("2: Fahrenheit to Celsius")
		fmt.Println("3: Celsius to Kelvin")
		fmt.Println("4: Kelvin to Celsius")
		fmt.Println("5: Meter to Foot")
		fmt.Println("6: Foot to Meter")
		fmt.Println("7: Pound to Kilogram")
		fmt.Println("8: Kilogram to Pound")
		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 1 and 8.")
			os.Exit(1)
		}
		switch choice {
		case 1:
			fmt.Printf("Enter temperature in Celsius: ")
			var c float64
			_, err := fmt.Scanln(&c)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				os.Exit(1)
			}
			f := conv.CToF(conv.Celsius(c))
			fmt.Printf("%s = %s\n", conv.Celsius(c), f)
		case 2:
			fmt.Printf("Enter temperature in Fahrenheit: ")
			var f float64
			_, err := fmt.Scanln(&f)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				os.Exit(1)
			}
			c := conv.FToC(conv.Fahrenheit(f))
			fmt.Printf("%s = %s\n", conv.Fahrenheit(f), c)
		case 3:
			fmt.Printf("Enter temperature in Celsius: ")
			var c float64
			_, err := fmt.Scanln(&c)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				os.Exit(1)
			}
			k := conv.CToK(conv.Celsius(c))
			fmt.Printf("%s = %s\n", conv.Celsius(c), k)
		case 4:
			fmt.Printf("Enter temperature in Kelvin: ")
			var k float64
			_, err := fmt.Scanln(&k)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				os.Exit(1)
			}
			c := conv.KtoC(conv.Kelvin(k))
			fmt.Printf("%s = %s\n", conv.Kelvin(k), c)
		case 5:
			fmt.Printf("Enter length in meters: ")
			var m float64
			_, err := fmt.Scanln(&m)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				os.Exit(1)
			}
			ft := conv.MtoFt(conv.Meter(m))
			fmt.Printf("%s = %s\n", conv.Meter(m), ft)
		case 6:
			fmt.Printf("Enter length in feet: ")
			var ft float64
			_, err := fmt.Scanln(&ft)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				os.Exit(1)
			}
			m := conv.FttoM(conv.Foot(ft))
			fmt.Printf("%s = %s\n", conv.Foot(ft), m)
		case 7:
			fmt.Printf("Enter weight in pounds: ")
			var p float64
			_, err := fmt.Scanln(&p)
			if err != nil {
				fmt.Println("Invalid input. Pleas enter a valid number.")
				os.Exit(1)
			}
			kg := conv.LbtoKg(conv.Pound(p))
			fmt.Printf("%s = %s\n", conv.Pound(p), kg)
		case 8:
			fmt.Printf("Enter weight in kilograms: ")
			var kg float64
			_, err := fmt.Scanln(&kg)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				os.Exit(1)
			}
			p := conv.KgtoLb(conv.Kilogram(kg))
			fmt.Printf("%s = %s\n", conv.Kilogram(kg), p)
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 8.")
			continue
		}
		fmt.Println("Do you want to perform another conversion? (y/n)")
		var again string
		_, err = fmt.Scanln(&again)
		if err != nil || (again != "y" && again != "Y") {
			fmt.Println("Thank you for using the unit conversion tool!")
			break
		}
		if again == "y" || again == "Y" {
			continue
		}
	}
}
