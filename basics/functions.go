package basics

// Function or Global Variable or type variable if defined with a Capital letter: They will be available/accessible from outside the package else they are available only inside package

func visibleInsidePackage() {
	// This function is available only within the package since it starts with lowercase letter
	// It can be also called an "unexported" method since it's name does not begins with a capital letter
}

func VisibleOutsidePackage() {
	// This function is available outside the package since it starts with uppercase letter
	// It can be also called an "exported" method since it's name begins with a capital letter
}

// function returning a single value - string
func UnderstandFunctionReturnSingleValue() string {
	return "Ankit"
}

// function returning multiple values - string and int
func UnderstandFunctionReturnMultipleValues() (string, int) {
	return "Ankit", 10
}
