package oddOrEven

func EvenOrOdd(number int) string {
	remainder := number % 2
	if(remainder != 0) {
		return "Odd"
	} else  {
		return "Even"
	}
}
