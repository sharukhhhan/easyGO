package domains

import "time"

type User struct {
	Id          int       `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	BirthDate   time.Time `json:"birthDate"`
	PhoneNumber string    `json:"phoneNumber"`
	Password    string    `json:"password"`
}
