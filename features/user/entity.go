package user

import "time"

type UserCore struct {
	Id          int
	Name        string `validate:"required"`
	PhoneNumber string `validate:"required"`
	Email       string `validate:"required"`
	Password    string `validate:"required"`
	Token       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Business interface {
	Create(data UserCore) (UserCore, error)
	GetAll() []UserCore
	GetById(id int) (UserCore, error)
	Update(data UserCore) (UserCore, error)
	Delete(id int) error
	Login(UserCore) (UserCore, error)
}

type Data interface {
	Create(data UserCore) (UserCore, error)
	GetAll() []UserCore
	GetById(id int) (UserCore, error)
	Update(data UserCore) (UserCore, error)
	Delete(id int) error
	Login(UserCore) (UserCore, error)
}
