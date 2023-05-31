package tempconv

func Ctof(c Celsius) Fahrenheit {
	c.String()
	return Fahrenheit(c*9/5 + 32)
}

func Ftoc(f Fahrenheit) Celsius {
	f.String()
	return Celsius((f - 32) * 5 / 9)
}

func Ctok(c Celsius) Kelvin {
	c.String()
	return Kelvin(c + 273.15)
}

func Ktoc(k Kelvin) Celsius {
	k.String()
	return Celsius(k - 273.15)
}

func Ftok(f Fahrenheit) Kelvin {
	f.String()
	return Kelvin((f-32)*5/9 + 273.15)
}

func Ktof(k Kelvin) Fahrenheit {
	k.String()
	return Fahrenheit((k-273.15)*9/5 + 32)
}
