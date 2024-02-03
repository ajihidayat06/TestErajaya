package setup

import (
	"TestErajaya/repository"
	"gorm.io/gorm"
)

type RegisterRepositoryStruct struct {
	ProductRepository repository.ProductRepository
}

func RegisterRepository(database *gorm.DB) RegisterRepositoryStruct {
	return RegisterRepositoryStruct{
		ProductRepository: repository.NewProductRepository(database),
	}
}
