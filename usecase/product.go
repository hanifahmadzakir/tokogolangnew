package usecase

import (
	"tokogolangnew/entity"
	"tokogolangnew/repository"

	"github.com/jinzhu/copier"
)

type IProductUsecase interface {
	CreateProduct(Product entity.ProductRequest) (entity.Product, error)
	GetAllProduct() ([]entity.Product, error)
	GetProductById(id int) (entity.Product, error)
	UpdateProduct(ProductRequest entity.UpdateProductRequest, id int) (entity.ProductResponse, error)
	DeleteProduct(id int) error
	GetProductByCategory(category string) ([]entity.Product, error)
	GetProductByPrice(priceMin int, priceMax int) ([]entity.Product, error)
}

type ProductUsecase struct {
	ProductRepository repository.IProductRespository
}

func NewProductUsecase(ProductRepository repository.IProductRespository) *ProductUsecase {
	return &ProductUsecase{ProductRepository}
}

func (usecase ProductUsecase) CreateProduct(Product entity.ProductRequest) (entity.ProductResponse, error) {
	u := entity.Product{
		Name:        Product.Name,
		Photo:       Product.Photo,
		Price:       Product.Price,
		Category:    Product.Category,
		Description: Product.Description,
	}

	Products, err := usecase.ProductRepository.Store(u)

	if err != nil {
		return entity.ProductResponse{}, err
	}

	ProductRes := entity.ProductResponse{
		ID:          Products.ID,
		Name:        Products.Name,
		Photo:       Products.Photo,
		Price:       Products.Price,
		Category:    Products.Category,
		Description: Products.Description,
	}
	return ProductRes, nil
}

func (usecase ProductUsecase) GetAllProduct() ([]entity.ProductResponse, error) {
	Products, err := usecase.ProductRepository.GetAll()

	if err != nil {
		return nil, err
	}

	ProductRes := []entity.ProductResponse{}
	copier.Copy(&ProductRes, &Products)
	return ProductRes, nil
}

func (usecase ProductUsecase) GetProductById(id int) (entity.ProductResponse, error) {
	product, err := usecase.ProductRepository.FindById(id)
	if err != nil {
		return entity.ProductResponse{}, err
	}
	productRes := entity.ProductResponse{}
	copier.Copy(&productRes, &product)
	return productRes, nil
}

func (usecase ProductUsecase) GetProductByCategory(category string) ([]entity.ProductResponse, error) {
	product, err := usecase.ProductRepository.FindByCategory(category)
	if err != nil {
		return []entity.ProductResponse{}, err
	}
	productRes := []entity.ProductResponse{}
	copier.Copy(&productRes, &product)
	return productRes, nil
}

func (usecase ProductUsecase) GetProductByPrice(priceMin int, priceMax int) ([]entity.ProductResponse, error) {
	product, err := usecase.ProductRepository.FindByPrice(priceMin, priceMax)
	if err != nil {
		return []entity.ProductResponse{}, err
	}
	productRes := []entity.ProductResponse{}
	copier.Copy(&productRes, &product)
	return productRes, nil
}

func (usecase ProductUsecase) UpdateProduct(ProductRequest entity.UpdateProductRequest, id int) (entity.ProductResponse, error) {

	Product, err := usecase.ProductRepository.Update(entity.Product{
		ID:          id,
		Name:        ProductRequest.Name,
		Photo:       ProductRequest.Photo,
		Price:       ProductRequest.Price,
		Category:    ProductRequest.Category,
		Description: ProductRequest.Description,
	})

	if err != nil {
		return entity.ProductResponse{}, err
	}

	copier.CopyWithOption(&Product, &ProductRequest, copier.Option{IgnoreEmpty: true})

	Product, err = usecase.ProductRepository.Update(Product)

	ProductRes := entity.ProductResponse{}

	copier.Copy(&ProductRes, &Product)

	return ProductRes, nil
}

func (usecase ProductUsecase) DeleteProduct(id int) error {
	_, err := usecase.ProductRepository.FindById(id)
	if err != nil {
		return err
	}
	err = usecase.ProductRepository.Delete(id)
	return err
}
