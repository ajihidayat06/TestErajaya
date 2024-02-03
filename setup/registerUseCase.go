package setup

import (
	"TestErajaya/usecase"
	"gorm.io/gorm"
)

type RegisterUseCaseStruct struct {
	ProductUseCase usecase.ProductUseCase
}

func RegisterUseCase(repo *RegisterRepositoryStruct, db *gorm.DB) RegisterUseCaseStruct {
	return RegisterUseCaseStruct{
		ProductUseCase: usecase.NewProductUseCase(repo.ProductRepository, db),
	}
}
