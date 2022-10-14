package repository

import (
	"tokogolangnew/entity"

	"gorm.io/gorm"
)

type ICartRespository interface {
	Store(cart entity.Cart) (entity.Cart, error)
	GetAll() ([]entity.Cart, error)
	Update(cart entity.Cart) (entity.Cart, error)
	Delete(id int) error
	FindById(id int) (entity.Cart, error)
}

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db}
}

func (r CartRepository) Store(Cart entity.Cart) (entity.Cart, error) {
	if err := r.db.Debug().Create(&Cart).Error; err != nil {

	}
	return Cart, nil
}

func (r CartRepository) GetAll() ([]entity.Cart, error) {
	var Carts []entity.Cart
	if err := r.db.Debug().Find(&Carts).Error; err != nil {
		return nil, err
	}
	return Carts, nil
}

func (repo CartRepository) FindById(id int) (entity.Cart, error) {
	var Cart entity.Cart
	if err := repo.db.Debug().Where("id = ?", id).First(&Cart).Error; err != nil {
		return entity.Cart{}, err
	}
	return Cart, nil
}

func (repo CartRepository) Update(Cart entity.Cart) (entity.Cart, error) {
	if err := repo.db.Debug().Save(&Cart).Error; err != nil {
		return entity.Cart{}, err
	}
	return Cart, nil
}

func (repo CartRepository) Delete(id int) error {
	if err := repo.db.Debug().Delete(&entity.Cart{}, id).Error; err != nil {
		return err
	}
	return nil
}
