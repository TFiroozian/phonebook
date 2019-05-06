package tmpl

type Contact struct {
	Id          int64  `json:"id" db:"id"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Email       string `json:"email" db:"email"`
}

type User struct {
	Id       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}
