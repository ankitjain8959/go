package basics

import "time"

// struct is a collection of fields and these fields are accessed using a dot
// Structure for User type
type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
	// Variables inside the type if starts with uppercase letter, are visible and accessible outside the package
	// Every variable has a **default value** (also known as the **zero value**) if not explicitly initialized.
}
