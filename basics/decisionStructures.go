package basics

import "log"

func UnderstandDecisionStructures() {
	// if else statement
	isHuman := false
	if isHuman {
		log.Println("I am a Human")
	} else {
		log.Println("I am not a Human")
	}

	//switch statement
	name := "Ankit"
	switch name {
	case "Ankit":
		log.Println("My Name is Ankit")
	case "Dimple":
		log.Println("My Name is Dimple")
	default:
		log.Println("My Name is something else")
	}
}