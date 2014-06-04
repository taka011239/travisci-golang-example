package fizzbuzz

import "strconv"

func fizzbuzz(v int) string {
	var fizz, buzz bool

	if v%3 == 0 {
		fizz = true
	}

	if v%5 == 0 {
		buzz = true
	}

	if fizz && buzz {
		return "fizzbuzz"
	} else if fizz {
		return "fizz"
	} else if buzz {
		return "buzz"
	} else {
		return strconv.Itoa(v)
	}
}
