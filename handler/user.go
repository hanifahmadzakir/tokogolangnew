package handler

import (
	"tokogolangnew/entity"
	"tokogolangnew/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserUsecase *usecase.UserUsecase
}

func NewUserHandler(UserUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{UserUsecase}
}

// CreateUsers godoc
// @Summary Create User.
// @Description create user.
// @Tags Users
// @Param Body body entity.UserRequest true "silahkan daftarkan data user"
// @Produce json
// @Success 201 {object} entity.UserResponse
// @Router /users/create [post]
func (handler UserHandler) CreateUser(c echo.Context) error {
	req := entity.UserRequest{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	user, err := handler.UserUsecase.CreateUser(req)

	if err != nil {
		return err
	}

	return c.JSON(201, user)
}

// GetAllUsers godoc
// @Summary Get All User.
// @Description get user.
// @Tags Users
// @Produce json
// @Success 201 {object} []entity.UserResponse
// @Router /users [get]
func (handler UserHandler) GetAllUser(c echo.Context) error {

	user, err := handler.UserUsecase.GetAllUser()

	if err != nil {
		return err
	}

	return c.JSON(200, user)
}

// GetUserById godoc
// @Summary Get User By Id
// @Description get user by id.
// @Tags Users
// @Produce json
// @Param id path string true "user id"
// @Success 200 {object} entity.UserResponse
// @Router /users/{id} [get]
func (handler UserHandler) GetUserByID(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))

	user, err := handler.UserUsecase.GetUserById(userId)
	if err != nil {
		return err
	}

	return c.JSON(200, user)
}

// UpdateUsers godoc
// @Summary Update Users
// @Description update users.
// @Tags Users
// @Param id path string true "user id"
// @Param Body body entity.UserRequest true "silahkan ubah data user"
// @Produce json
// @Success 200 {object} entity.UserResponse
// @Router /users/update/{id} [put]
func (handler UserHandler) UpdateUser(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	userRequest := entity.UpdateUserRequest{}
	if err := c.Bind(&userRequest); err != nil {
		return err
	}
	user, err := handler.UserUsecase.UpdateUser(userRequest, userId)
	if err != nil {
		return err
	}

	return c.JSON(200, user)
}

// DeleteUser godoc
// @Summary Delete User
// @Description Delete User by id.
// @Tags Users
// @Param id path string true "user id"
// @Produce json
// @Success 200 {object} string
// @Router /users/delete/{id} [delete]
func (handler UserHandler) DeleteUser(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))
	err := handler.UserUsecase.DeleteUser(userId)
	if err != nil {
		return err
	}
	return c.JSON(200, "user berhasil di delete")
}
