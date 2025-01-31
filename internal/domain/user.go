package domain

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	OptedOut *bool  `json:"opted_out"`
}
