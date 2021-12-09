package mygo

import (
	"fmt"
	"math/rand"
)

func TestSlices() {
	var arr = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr)

	slice1 := arr[0:2]
	fmt.Println("Slice1:", slice1)
	slice2 := arr[:]
	fmt.Println("Slice2:", slice2)
	slice3 := arr[:5]
	fmt.Println("Slice3:", slice3)

	slice4 := make([]int, 0)
	for i := 0; i < 11; i++ {
		slice4 = append(slice4, i)
	}
	fmt.Println("Slice4:", slice4)

	slice5 := make([]int, 5, 10)
	slice5[3] = 3 * 3
	fmt.Println("Slice5 len:", len(slice5), "cap:", cap(slice5))
	fmt.Println(slice5)
	for i := 0; i < 11; i++ {
		slice5 = append(slice5, i)
	}
	fmt.Println("Slice5 after append:", slice5)

	appendRandToSlice()

}

func appendRandToSlice() {
	slice1 := make([]int, 0)
	for i := 0; i < 100; i++ {
		slice1 = append(slice1, rand.Intn(100))
	}
	for i := 0; i < 100; i++ {
		//slice1 = append(slice1, rand.Intn(100))
		print(slice1[i], ", ")
	}
	//println("Slice with random values:", slice1)
}
