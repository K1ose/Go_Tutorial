package weightconv

func Kgtojin(k Kilo) Jin {
	return Jin(k * 2)
}

func Jintokg(j Jin) Kilo {
	return Kilo(j / 2)
}

func Poundtokg(p Pound) Kilo {
	return Kilo(p * 0.4536)
}

func Kgtopound(k Kilo) Pound {
	return Pound(k / 0.4536)
}

func Jintopound(j Jin) Pound {
	return Pound(j / 0.9072)
}

func Poundtojin(p Pound) Jin {
	return Jin(p * 0.9072)
}
