package handler


import (
	"tokogolangnew/entity"
	"tokogolangnew/usecase"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	ProductUsecase *usecase.ProductUsecase
}

func NewProductHandler(ProductUsecase *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{ProductUsecase}
}

// Create products godoc
// @Summary Create Product.
// @Description create product.
// @Tags Products
// @Param Body formData entity.ProductRequest true "silahkan daftarkan data product"
// @Consumes multipart/form-data
// @Produce json
// @Success 201 {object} entity.ProductResponse
// @Router /products/create [post]
func (handler ProductHandler) CreateProduct(c echo.Context) error {
	name := c.FormValue("name")
	price, _ := strconv.Atoi(c.FormValue("price"))
	category := c.FormValue("category")
	description := c.FormValue("description")
	file, err := c.FormFile("photo")

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	// destinattion
	dst, err := os.Create(fmt.Sprintf("%s%s", "storage/", file.Filename))
	if err != nil {
		return err
	}

	defer dst.Close()

	// Copy
	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/%s", os.Getenv("BASE_URL"), dst.Name())

	req := entity.ProductRequest{
		Name:        name,
		Photo:       filePath,
		Price:       price,
		Category:    category,
		Description: description,
	}
	Product, err := handler.ProductUsecase.CreateProduct(req)

	if err != nil {
		return err
	}

	return c.JSON(201, Product)
}

// GetAllProducts godoc
// @Summary Get All Product.
// @Description get product.
// @Tags Products
// @Produce json
// @Success 201 {object} []entity.ProductResponse
// @Router /products [get]
func (handler ProductHandler) GetAllProduct(c echo.Context) error {

	Product, err := handler.ProductUsecase.GetAllProduct()

	if err != nil {
		return err
	}

	return c.JSON(200, Product)
}

// GetUserById godoc
// @Summary Get User By Id
// @Description get User by id.
// @Tags Products
// @Produce json
// @Param id path string true "user id"
// @Success 200 {object} entity.ProductResponse
// @Router /products/{id} [get]
func (handler ProductHandler) GetProductByID(c echo.Context) error {
	productId, _ := strconv.Atoi(c.Param("id"))

	product, err := handler.ProductUsecase.GetProductById(productId)
	if err != nil {
		return err
	}

	return c.JSON(200, product)
}

// Get Products by category godoc
// @Summary Get Products by category
// @Description get products by category.
// @Tags Products
// @Produce json
// @Param category path string true "product category"
// @Success 200 {object} string
// @Router /products/{category} [get]
func (handler ProductHandler) GetProductByCategory(c echo.Context) error {
	productCategory := c.Param("category")

	product, err := handler.ProductUsecase.GetProductByCategory(productCategory)
	if err != nil {
		return err
	}

	return c.JSON(200, product)
}

// Get Products by price godoc
// @Summary Get Products by price
// @Description Get Products by price.
// @Tags Products
// @Produce json
// @Param Body priceMin,priceMax path string true "product price"
// @Success 200 {object} string
// @Router /products/{priceMin}/{priceMax} [get]
func (handler ProductHandler) GetProductByPrice(c echo.Context) error {
	priceMin, _ := strconv.Atoi(c.Param("priceMin"))
	priceMax, _ := strconv.Atoi(c.Param("priceMax"))

	product, err := handler.ProductUsecase.GetProductByPrice(priceMin, priceMax)
	if err != nil {
		return err
	}
	return c.JSON(200, product)
}

// Update products godoc
// @Summary Update Product.
// @Description Update product.
// @Tags Products
// @Param Body formData entity.ProductRequest true "silahkan update data product"
// @Consumes multipart/form-data
// @Produce json
// @Success 201 {object} entity.ProductResponse
// @Router /products/update [post]
func (handler ProductHandler) UpdateProduct(c echo.Context) error {
	ProductId, err := strconv.Atoi(c.Param("id"))
	name := c.FormValue("name")
	price, _ := strconv.Atoi(c.FormValue("price"))
	category := c.FormValue("category")
	description := c.FormValue("description")
	file, err := c.FormFile("photo")
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	// destinattion
	dst, err := os.Create(fmt.Sprintf("%s%s", "storage/", file.Filename))
	if err != nil {
		return err
	}

	defer dst.Close()

	// Copy
	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/%s", os.Getenv("BASE_URL"), dst.Name())

	req := entity.UpdateProductRequest{
		Name:        name,
		Photo:       filePath,
		Price:       price,
		Category:    category,
		Description: description,
	}
	ProductRequest := entity.UpdateProductRequest{}
	if err := c.Bind(&ProductRequest); err != nil {
		return err
	}
	Product, err := handler.ProductUsecase.UpdateProduct(req, ProductId)
	if err != nil {
		return err
	}

	return c.JSON(200, Product)
}

// DeleteProduct godoc
// @Summary Delete Product
// @Description Delete Product by id.
// @Tags Products
// @Produce json
// @Param id path int true "product id"
// @Success 200 {object} string
// @Router /products/{id} [delete]
func (handler ProductHandler) DeleteProduct(c echo.Context) error {
	ProductId, _ := strconv.Atoi(c.Param("id"))
	err := handler.ProductUsecase.DeleteProduct(ProductId)
	if err != nil {
		return err
	}
	return c.JSON(200, ProductId)
}