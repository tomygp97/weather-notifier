package domain

type UserRepository interface {
	Save(user *User) error
	FindByID(id int) (*User, error)
	FindAll() ([]User, error)
	Update(user *User) error
	Delete(id int) error
}
