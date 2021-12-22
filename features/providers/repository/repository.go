package repository

import (
	"bayareen-backend/features/providers"

	"gorm.io/gorm"
)

type posgresRepository struct {
	Conn *gorm.DB
}

func NewPostgresRepository(conn *gorm.DB) providers.Data {
	return &posgresRepository{
		Conn: conn,
	}
}

func (repo *posgresRepository) Create(data *providers.Core) (*providers.Core, error) {
	record := FromCore(data)

	if err := repo.Conn.Create(record).Error; err != nil {
		return &providers.Core{}, err
	}

	return record.ToCore(), nil
}

func (repo *posgresRepository) GetAll() []providers.Core {
	records := []Provider{}
	repo.Conn.Find(&records)

	return ToCoreSlice(records)
}

func (repo *posgresRepository) GetById(id int) (*providers.Core, error) {
	record := Provider{Id: id}
	if err := repo.Conn.First(&record).Error; err != nil {
		return &providers.Core{}, err
	}

	return record.ToCore(), nil
}