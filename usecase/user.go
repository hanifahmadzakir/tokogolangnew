package usecase

import (
	"tokogolangnew/entity"
	"tokogolangnew/repository"

	"github.com/jinzhu/copier"
)

type IUserUsecase interface {
	CreateUser(user entity.UserRequest) (entity.User, error)
	GetAllUser() ([]entity.User, error)
	GetProductById(id int) (entity.Product, error)
	UpdateUser(userRequest entity.UpdateUserRequest, id int) (entity.UserResponse, error)
	DeleteUser(id int) error
}

type UserUsecase struct {
	userRepository repository.IUserRespository
}

func NewUserUsecase(userRepository repository.IUserRespository) *UserUsecase {
	return &UserUsecase{userRepository}
}

func (usecase UserUsecase) CreateUser(user entity.UserRequest) (entity.UserResponse, error) {
	u := entity.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	users, err := usecase.userRepository.Store(u)

	if err != nil {
		return entity.UserResponse{}, err
	}

	userRes := entity.UserResponse{
		ID:       users.ID,
		Name:     users.Name,
		Email:    users.Email,
		Password: users.Password,
	}
	return userRes, nil
}

func (usecase UserUsecase) GetAllUser() ([]entity.UserResponse, error) {
	users, err := usecase.userRepository.GetAll()

	if err != nil {
		return nil, err
	}

	userRes := []entity.UserResponse{}
	copier.Copy(&userRes, &users)
	return userRes, nil
}

func (usecase UserUsecase) GetUserById(id int) (entity.UserResponse, error) {
	user, err := usecase.userRepository.FindById(id)
	if err != nil {
		return entity.UserResponse{}, err
	}
	userRes := entity.UserResponse{}
	copier.Copy(&userRes, &user)
	return userRes, nil
}

func (usecase UserUsecase) UpdateUser(userRequest entity.UpdateUserRequest, id int) (entity.UserResponse, error) {

	user, err := usecase.userRepository.Update(entity.User{
		ID:       id,
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	})

	if err != nil {
		return entity.UserResponse{}, err
	}

	copier.CopyWithOption(&user, &userRequest, copier.Option{IgnoreEmpty: true})

	user, err = usecase.userRepository.Update(user)

	userRes := entity.UserResponse{}

	copier.Copy(&userRes, &user)

	return userRes, nil
}

func (usecase UserUsecase) DeleteUser(id int) error {
	_, err := usecase.userRepository.FindById(id)
	if err != nil {
		return err
	}
	err = usecase.userRepository.Delete(id)
	return err
}
