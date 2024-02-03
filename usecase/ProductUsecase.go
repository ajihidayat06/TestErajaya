package usecase

import (
	"TestErajaya/model"
	"TestErajaya/repository"
	"gorm.io/gorm"
)

type ProductUseCase interface {
	Create(request model.Product) (response interface{}, err error)
	List(userParam model.ListDataRequestStruct) (responses []model.ProductResponse)
}

type productUseCaseImpl struct {
	ProductRepository repository.ProductRepository
	DB                *gorm.DB
}

func NewProductUseCase(productRepository repository.ProductRepository, db *gorm.DB) ProductUseCase {
	return &productUseCaseImpl{
		ProductRepository: productRepository,
		DB:                db,
	}
}

func (p *productUseCaseImpl) Create(product model.Product) (response interface{}, err error) {
	response, err = InsertWithTx(p.DB, p.doInsert, product)
	if err != nil {
		return
	}

	return
}

func (p *productUseCaseImpl) doInsert(tx *gorm.DB, modelData interface{}) (response interface{}, err error) {
	product := modelData.(model.Product)

	err = p.ProductRepository.Insert(tx, product)
	if err != nil {
		return
	}

	return
}

func (p *productUseCaseImpl) List(userParam model.ListDataRequestStruct) (responses []model.ProductResponse) {
	products := p.ProductRepository.FindAll(userParam)
	responses = reformatToResposeProducts(products)
	return
}

func reformatToResposeProducts(products []model.Product) (results []model.ProductResponse) {
	for i := 0; i < len(products); i++ {
		product := model.ProductResponse{
			Id:          products[i].Id,
			Name:        products[i].Name,
			Price:       products[i].Price,
			Description: products[i].Description,
			Quantity:    products[i].Quantity,
		}

		results = append(results, product)
	}

	return
}
