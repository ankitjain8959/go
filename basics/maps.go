package basics

import "log"

func UnderstandMap() {
	// Empty map or nil map.
	// Zero value of map is nil.
	// A nil map doesn't contain any key. If you try to add a key-value pair in the nil map, then compiler will throw runtime error
	var nilMap map[int]string
	if nilMap == nil {
		log.Println("Empty Map")
	}

	// Creating and initializing a map using a short hand declaration
	firstNamesMap := map[int]string{
		1: "Ankit",
		2: "Dimple",
	}
	log.Println("FirstNames Map values are:", firstNamesMap[1], "-", firstNamesMap[2])

	// Creating a map using make() function. make() function always returns a map which is initialized
	lastNamesMap := make(map[int]string)
	lastNamesMap[1] = "Jain"
	lastNamesMap[2] = "Mehta"

	log.Println("LastNames Map values are:", lastNamesMap[1], "-", lastNamesMap[2])

	customMap := make(map[int]User)
	myUser := User{
		FirstName: "Ankit",
		LastName:  "Jain",
	}
	customMap[1] = myUser
	log.Println("My Custom Map values are", customMap[1])
	log.Println("My Custom Map FirstName value is", customMap[1].FirstName)
}
