package basics

import "time"

// struct is a collection of fields and these fields are accessed using a dot
// Structure for User type
type User struct {
	FirstName   string    `json:"firstName,bson:"name"`
	LastName    string    `json:"lastName,omitempty,bson:"name,omitempty"`
	PhoneNumber string    `json:"phoneNumber"`
	Age         int       `json:"age,omitempty"`
	BirthDate   time.Time `json:"birthDate"` // time.Time is a built-in type in Go for handling date and time

	// Variables inside the type if starts with uppercase letter, are visible and accessible outside the package
	// Every variable has a **default value** (also known as the **zero value**) if not explicitly initialized.

	// The struct tags (e.g., json:"statusCode" or bson:"statusCode,omitempty") are used to control how the struct fields are serialized to JSON or other formats.
	// Using json:"statusCode" tag specify how the field will be named when the struct is converted to JSON. For example, StatusCode will appear as statusCode in the JSON output.
	// Using bson:"statusCode,omitempty" tag specify how the field will be named when the struct is converted to BSON (Binary JSON, used in MongoDB). The omitempty option means that if the field has a zero value (like an empty string or zero integer), it will be omitted from the BSON output.
	// Using omitempty tags allows omitting the field from JSON if it has a zero value.
}
