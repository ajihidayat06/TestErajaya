package repository

import (
	"TestErajaya/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Insert(tx *gorm.DB, product model.Product) error
	FindAll(userParam model.ListDataRequestStruct) (products []model.Product)
}

type productRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(database *gorm.DB) ProductRepository {
	return &productRepositoryImpl{
		DB: database,
	}
}

func (r *productRepositoryImpl) Insert(tx *gorm.DB, product model.Product) error {
	result := tx.Create(&product)
	return result.Error
}

func (r *productRepositoryImpl) FindAll(userParam model.ListDataRequestStruct) (products []model.Product) {
	db := convertUserParamToDBQuery(r.DB, userParam)
	db.Find(&products)
	return
}
