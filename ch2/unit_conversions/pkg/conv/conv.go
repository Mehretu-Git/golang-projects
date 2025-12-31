package conv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Meter float64
type Foot float64
type Pound float64
type Kilogram float64

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }
func (m Meter) String() string      { return fmt.Sprintf("%gm", m) }
func (ft Foot) String() string      { return fmt.Sprintf("%gft", ft) }
func (p Pound) String() string      { return fmt.Sprintf("%glb", p) }
func (kg Kilogram) String() string  { return fmt.Sprintf("%gkg", kg) }

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius(5 * (f - 32) / 9)
}
func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}
func KtoC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}
func MtoFt(m Meter) Foot {
	return Foot(m * 3.28084)
}
func FttoM(ft Foot) Meter {
	return Meter(ft / 3.28084)
}
func LbtoKg(p Pound) Kilogram {
	return Kilogram(p / 2.20462)
}
func KgtoLb(kg Kilogram) Pound {
	return Pound(kg * 2.20462)
}
