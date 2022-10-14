package repository

import (
	"tokogolangnew/entity"

	"gorm.io/gorm"
)

type IProductRespository interface {
	Store(product entity.Product) (entity.Product, error)
	GetAll() ([]entity.Product, error)
	Update(product entity.Product) (entity.Product, error)
	Delete(id int) error
	FindById(id int) (entity.Product, error)
	FindByCategory(category string) ([]entity.Product, error)
	FindByPrice(priceMin int, priceMax int) ([]entity.Product, error)
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (r ProductRepository) Store(product entity.Product) (entity.Product, error) {
	if err := r.db.Debug().Create(&product).Error; err != nil {

	}
	return product, nil
}

func (r ProductRepository) GetAll() ([]entity.Product, error) {
	var Products []entity.Product
	if err := r.db.Debug().Find(&Products).Error; err != nil {
		return nil, err
	}
	return Products, nil
}

func (repo ProductRepository) FindById(id int) (entity.Product, error) {
	var Product entity.Product
	if err := repo.db.Debug().Where("id = ?", id).First(&Product).Error; err != nil {
		return entity.Product{}, err
	}
	return Product, nil
}

func (r ProductRepository) FindByCategory(category string) ([]entity.Product, error) {
	var product []entity.Product
	if err := r.db.Debug().Where("category = ?", category).Find(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r ProductRepository) FindByPrice(priceMin int, priceMax int) ([]entity.Product, error) {
	var product []entity.Product
	if err := r.db.Debug().Where("price >= ? AND price <= ? ", priceMin, priceMax).Find(&product).Error; err != nil {
		return []entity.Product{}, err
	}
	return product, nil
}

func (repo ProductRepository) Update(product entity.Product) (entity.Product, error) {
	if err := repo.db.Debug().Save(&product).Error; err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

func (repo ProductRepository) Delete(id int) error {
	if err := repo.db.Debug().Delete(&entity.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}
