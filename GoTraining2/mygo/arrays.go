package mygo

func TestArrays() {

	arr1 := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		println(arr1[i])
	}

	arr2 := [7]int{1, 2, 3, 4, 5, 6, 7}
	for i := range arr2 {
		println(arr2[i])
	}

	arr3 := [2][2]string{{"srini", "edara"}, {"Rao", "ramela"}}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			println(arr3[i][j])
		}
	}

}
