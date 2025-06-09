package basics

import "log"

// Go language contains only one loop i.e. for-loop. However, the for loop can be used in different form
func UnderstandLoops() {
	// 1. simple for loop with it's three component: the init statement, the condition expression and the post statement
	for index := 0; index < 10; index++ {
		log.Println("simple for loop with it's three component", index)
	}

	animals := []string{"Dog", "Cat"}
	for index := 0; index < len(animals); index++ {
		log.Println("simple for loop with it's three component", animals[index])
	}

	// init statement and post statement are optional
	// 2. In Go, for loop can work as a while loop
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	log.Println("for loop works as while loop", sum)

	// 3.1. using range in for loop
	for index, animal := range animals {
		log.Println("using range in for loops", index, animal)
	}

	// 3.2. using range in for loops and using underscore when you don't care about the index/current iteration
	for _, animal := range animals {
		log.Println("using range in for loops with underscore in place of index", animal)
	}

	//Note: we can range over custom objects, slices, string, maps

	// 4. for loop can iterate over the key and value pairs of the map
	myMap := map[int]string{
		1: "Ankit",
		2: "Dimple",
	}
	for key, value := range myMap {
		log.Println("Iterate map using for loop:", key, value)
	}
	// 5. infinite for loop when all three component are missing: the init statement, the condition expression and the post statement
	/*
		for {}
	*/
}
