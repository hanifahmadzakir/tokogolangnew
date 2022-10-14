package usecase

import (
	"errors"
	"tokogolangnew/config"
	"tokogolangnew/entity"
	"tokogolangnew/repository"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type ICartUsecase interface {
	CreateCart(Cart entity.CartRequest) (entity.Cart, error)
	SelectAllCart() ([]entity.Cart, error)
	GetProductById(id int) (entity.Product, error)
	UpdateCart(CartRequest entity.UpdateCartRequest, id int) (entity.CartResponse, error)
	DeleteCart(id int) error
}

type CartUsecase struct {
	CartRepository repository.ICartRespository
}

func NewCartUsecase(CartRepository repository.ICartRespository) *CartUsecase {
	return &CartUsecase{CartRepository}
}

func getUserByID(e int) (entity.User, error) {
	var user entity.User

	if err := config.DB.Model(entity.User{}).Where("id = ?", e).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, nil
		}
		return user, err
	}
	return user, nil
}

func getProductByID(e int) (entity.Product, error) {
	var product entity.Product

	if err := config.DB.Model(entity.Product{}).Where("id = ?", e).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return product, nil
		}
		return product, err
	}
	return product, nil
}

func (usecase CartUsecase) CreateCart(Cart entity.CartRequest) (entity.CartResponse, error) {
	u := entity.Cart{
		Quantity:  Cart.Quantity,
		Checkout:  Cart.Checkout,
		UserID:    Cart.UserID,
		ProductID: Cart.ProductID,
	}

	Carts, err := usecase.CartRepository.Store(u)

	if err != nil {
		return entity.CartResponse{}, err
	}
	user, _ := getUserByID(Cart.UserID)
	product, _ := getProductByID(Cart.ProductID)

	CartRes := entity.CartResponse{
		ID:        Carts.ID,
		Quantity:  Carts.Quantity,
		Checkout:  Carts.Checkout,
		UserID:    Carts.UserID,
		User:      user,
		ProductID: Carts.ProductID,
		Product:   product,
	}
	return CartRes, nil
}

func (usecase CartUsecase) SelectAllCart() ([]entity.CartResponse, error) {
	Carts, err := usecase.CartRepository.GetAll()

	if err != nil {
		return nil, err
	}

	CartRes := []entity.CartResponse{}
	copier.Copy(&CartRes, &Carts)
	return CartRes, nil
}

func (usecase CartUsecase) GetCartById(id int) (entity.CartResponse, error) {
	Cart, err := usecase.CartRepository.FindById(id)
	if err != nil {
		return entity.CartResponse{}, err
	}
	CartRes := entity.CartResponse{}
	copier.Copy(&CartRes, &Cart)
	return CartRes, nil
}

func (usecase CartUsecase) UpdateCart(CartRequest entity.UpdateCartRequest, id int) (entity.CartResponse, error) {

	Cart, err := usecase.CartRepository.Update(entity.Cart{
		ID:        id,
		Quantity:  CartRequest.Quantity,
		Checkout:  CartRequest.Checkout,
		UserID:    CartRequest.UserID,
		ProductID: CartRequest.ProductID,
	})

	if err != nil {
		return entity.CartResponse{}, err
	}

	copier.CopyWithOption(&Cart, &CartRequest, copier.Option{IgnoreEmpty: true})

	Cart, err = usecase.CartRepository.Update(Cart)

	CartRes := entity.CartResponse{}

	copier.Copy(&CartRes, &Cart)

	return CartRes, nil
}

func (usecase CartUsecase) DeleteCart(id int) error {
	_, err := usecase.CartRepository.FindById(id)
	if err != nil {
		return err
	}
	err = usecase.CartRepository.Delete(id)
	return err
}
