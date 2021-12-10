package mygo

import "fmt"

func TestEmpyInterfaces() {

	fmt.Println("Add1: ", add(10, 20))
	fmt.Println("Add1: ", add(float32(10.10), float32(20.40)))
	fmt.Println("Add1: ", add(10, float32(20.10)))

}

func add(values ...interface{}) (sum interface{}) {

	for _, value := range values {

		switch value.(type) {

		case int:
			if sum == nil {
				sum = int(0)
			}
			sum = sum.(int) + value.(int)

		case float32:
			if sum == nil {
				sum = float32(0)
			}
			sum = sum.(float32) + value.(float32)
		}
	}
	return sum
}
