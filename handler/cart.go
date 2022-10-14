package handler

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"tokogolangnew/config"
	"tokogolangnew/entity"
	"tokogolangnew/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CartHandler struct {
	CartUsecase *usecase.CartUsecase
}

func NewCartHandler(CartUsecase *usecase.CartUsecase) *CartHandler {
	return &CartHandler{CartUsecase}
}

// Create Cart godoc
// @Summary Create Cart.
// @Description Create Cart.
// @Tags Carts
// @Param Body body entity.CartResponse true "Buat cart"
// @Produce json
// @Success 200 {object} entity.CartResponse
// @Router /carts [post]
func (handler CartHandler) CreateCart(c echo.Context) error {
	req := entity.CartRequest{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	Cart, err := handler.CartUsecase.CreateCart(req)

	if err != nil {
		return err
	}

	return c.JSON(201, Cart)
}

// Get All Carts godoc
// @Summary Get All Carts.
// @Description get All Carts.
// @Tags Carts
// @Produce json
// @Success 200 {object} entity.CartResponse
// @Router /carts [get]
func (handler CartHandler) SelectAllCart(c echo.Context) error {

	Cart, err := handler.CartUsecase.SelectAllCart()

	if err != nil {
		return err
	}

	return c.JSON(200, Cart)
}

func (handler CartHandler) GetCartByID(c echo.Context) error {
	CartId, _ := strconv.Atoi(c.Param("id"))

	Cart, err := handler.CartUsecase.GetCartById(CartId)
	if err != nil {
		return err
	}

	return c.JSON(200, Cart)
}

func (handler CartHandler) UpdateCart(c echo.Context) error {
	CartId, err := strconv.Atoi(c.Param("id"))
	CartRequest := entity.UpdateCartRequest{}
	if err := c.Bind(&CartRequest); err != nil {
		return err
	}
	Cart, err := handler.CartUsecase.UpdateCart(CartRequest, CartId)
	if err != nil {
		return err
	}

	return c.JSON(200, Cart)
}

func (handler CartHandler) DeleteCart(c echo.Context) error {
	CartId, _ := strconv.Atoi(c.Param("id"))
	err := handler.CartUsecase.DeleteCart(CartId)
	if err != nil {
		return err
	}
	return c.JSON(200, CartId)
}

func (handler CartHandler) Payment(c echo.Context) error {

	userId, _ := strconv.Atoi(c.Param("id"))
	file, err := c.FormFile("payment")

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	// destinattion
	dst, err := os.Create(fmt.Sprintf("%s%s", "/storage", file.Filename))
	if err != nil {
		return err
	}

	defer dst.Close()

	// Copy
	if _, err := io.Copy(dst, src); err != nil {
		return err
	}
	filePath := fmt.Sprintf("%s/%s", os.Getenv("BASE_URL"), dst.Name())
	if payment := config.DB.Model(entity.Cart{}).Where("user_id = ?", userId).Update("checkout", true).Error; err != nil {
		if errors.Is(payment, gorm.ErrRecordNotFound) {
			return err
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Payment",
		"data":    filePath,
	})
}
