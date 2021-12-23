package repository

import (
	"bayareen-backend/features/paymentmethods"

	"gorm.io/gorm"
)

type posgresPaymentMethodRepository struct {
	Conn *gorm.DB
}

func NewPostgresPaymentMethodRepository(conn *gorm.DB) paymentmethods.Data {
	return &posgresPaymentMethodRepository{
		Conn: conn,
	}
}

func (repo *posgresPaymentMethodRepository) Create(data *paymentmethods.Core) (*paymentmethods.Core, error) {
	record := FromCore(data)
	if err := repo.Conn.Create(record).Error; err != nil {
		return &paymentmethods.Core{}, err
	}

	return record.ToCore(), nil
}

func (repo *posgresPaymentMethodRepository) GetAll() []paymentmethods.Core {
	records := []PaymentMethod{}
	repo.Conn.Find(&records)

	return ToCoreSlice(records)
}

func (repo *posgresPaymentMethodRepository) GetById(id int) (*paymentmethods.Core, error) {
	record := PaymentMethod{
		Id: id,
	}
	if err := repo.Conn.First(&record).Error; err != nil {
		return &paymentmethods.Core{}, err
	}

	return record.ToCore(), nil
}

func (repo *posgresPaymentMethodRepository) Update(data *paymentmethods.Core) (*paymentmethods.Core, error) {
	record := FromCore(data)
	if err := repo.Conn.Save(&record).Error; err != nil {
		return &paymentmethods.Core{}, err
	}

	return record.ToCore(), nil
}
