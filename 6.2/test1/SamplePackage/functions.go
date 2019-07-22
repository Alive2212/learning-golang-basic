package main

// List of functions
func NewUser(firstName, lastName string) *User {
	return &User{FirstName: firstName,
		LastName: lastName,
		Location: &UserLocation{
			City:
			"Santa Monica",
			Country: "USA",
		},
	}
}

