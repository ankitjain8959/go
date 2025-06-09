package basics

import (
	"log"
	"sort"
)

func UnderstandSlices() {

	//slice (like, arrays) of integers
	var numbers []int
	numbers = append(numbers, 3)
	numbers = append(numbers, 1)
	numbers = append(numbers, 2)
	log.Println("Slice values are", numbers)

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	log.Println("Slice values after sorting in ascending order are", numbers)

	// Zero value of slice is nil.
	// a nil slice has a length and capacity of 0
	var nilSlice []int
	log.Println("Nil Slice is ", nilSlice) //Output: "Nil Slice is []"

	newNumbers := []int{5, 3, 2, 4, 1}
	log.Println("New Slice values are", newNumbers)

	//Range of values from the slice - newNumber[x:y] where x is inclusive and y is exclusive
	log.Println(newNumbers[0:2])
	log.Println(newNumbers[1:3])
}
