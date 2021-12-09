package mygo

func TestConditionals() {

	isValid := false

	if isValid {
		return
	}

	num := 21
	if num%2 == 0 {
		println("Its even number")
	} else {
		println("Its odd number")
	}
}
