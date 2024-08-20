package main

func conditionWithIf() {
	var currentYear = 2021

	if age := currentYear - 1999; age < 17 {
		println("You are not allowed to drive")
	} else if age >= 17 && age <= 18 {
		println("You are allowed to drive with supervision")
	} else {
		println("You are allowed to drive")
	}
}

func conditionWithSwitch() {
	var score = 7

	switch score {
	case 10:
		println("Perfect!")
	case 8, 9:
		println("Awesome!")
	case 6, 7:
		{
			println("Good!")
		}
	default:
		println("Not bad!")
	}
}

func conditionWithSwitchFallthrough() {
	var score = 5

	switch {
	case score == 10:
		println("Perfect!")
	case score >= 8 && score <= 9:
		println("Awesome!")

	case score >= 6 && score <= 7:
		{
			println("Good!")

		}
	case score == 5:
		println("Not bad!")
		fallthrough
	case score <= 5:
		println("You need to learn more!")
	default:
		println("Not bad!")
	}
}

func main() {
	conditionWithSwitchFallthrough()
}
