package admins

import "time"

type Core struct {
	Id        int
	Name      string `validate:"required"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	Create(data *Core) (*Core, error)
	GetAll() []Core
	GetById(id int) (*Core, error)
	Update(data *Core) (*Core, error)
	Delete(id int) error
}

type Data interface {
	Create(data *Core) (*Core, error)
	GetAll() []Core
	GetById(id int) (*Core, error)
	Update(data *Core) (*Core, error)
	Delete(id int) error
}