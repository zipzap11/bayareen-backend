package repository

import (
	"bayareen-backend/features/categories"

	"gorm.io/gorm"
)

type postgreRepository struct {
	Conn *gorm.DB
}

func NewPostgreRepository(conn *gorm.DB) categories.Data {
	return &postgreRepository{
		Conn: conn,
	}
}

func (repo *postgreRepository) Create(core categories.Core) (categories.Core, error) {
	record := FromCore(core)
	if err := repo.Conn.Create(&record).Error; err != nil {
		return categories.Core{}, err
	}

	return record.ToCore(), nil
}

func (repo *postgreRepository) GetAll() []categories.Core {
	resp := []Category{}
	repo.Conn.Find(&resp)

	return ToCoreSlice(&resp)
}