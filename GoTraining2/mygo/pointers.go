package mygo

func TestPointers() {
	var ptr *int
	ptr = returnLocalAddress()

	println("Value=", *ptr)

}

func returnLocalAddress() *int {
	var i int = 10
	i++
	return &i
}
