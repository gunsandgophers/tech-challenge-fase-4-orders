package controllers

import (
	"net/http"
	"strconv"
	"tech-challenge-fase-1/internal/core/dtos"
	"tech-challenge-fase-1/internal/core/errors"
	"tech-challenge-fase-1/internal/core/repositories"
	"tech-challenge-fase-1/internal/core/use_cases/products"
	httpserver "tech-challenge-fase-1/internal/infra/http"
)

type ProductController struct {
	productRepository repositories.ProductRepositoryInterface
}

type ProductRequest struct {
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Image       string  `json:"image" `
}

func NewProductController(
	productRepository repositories.ProductRepositoryInterface,
) *ProductController {
	return &ProductController{
		productRepository: productRepository,
	}
}

// CreateProduct godoc
//
//	@Summary		Create a product
//	@Description	register the product information
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			product	body		ProductRequest	true	"Insert Product"
//	@Success		200		{object}	dtos.ProductDTO
//	@Failure		400		{string}	string	"when bad request"
//	@Failure		500		{string}	string	"when create product process error"
//	@Router			/product/ [post]
func (pc *ProductController) CreateProduct(c httpserver.HTTPContext) {
	var product ProductRequest
	c.BindJSON(&product)
	if err := product.ValidateProduct(); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	respProduct, err := products.NewCreateProductUseCase(pc.productRepository).Execute(&dtos.ProductDTO{
		ID:          "",
		Name:        product.Name,
		Category:    product.Category,
		Price:       product.Price,
		Description: product.Description,
		Image:       product.Image,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpserver.Payload{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, httpserver.Payload{
		"product": respProduct,
	})
}

// UpdateProduct godoc
//
//	@Summary		Update a product
//	@Description	change the product information
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string			true	"Product identifier"
//	@Param			product	body		ProductRequest	true	"Product object"
//	@Success		200		{object}	dtos.ProductDTO
//	@Failure		400		{string}	string	"when bad request"
//	@Failure		500		{string}	string	"when update product process error"
//	@Router			/product/{id}/ [put]
func (pc *ProductController) UpdateProduct(c httpserver.HTTPContext) {
	var product ProductRequest
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, httpserver.Payload{
			"error": err.Error(),
		})
		return
	}

	respProduct, err := products.NewUpdateProductUseCase(pc.productRepository).Execute(&dtos.ProductDTO{
		ID:          c.Param("id"),
		Name:        product.Name,
		Category:    product.Category,
		Price:       product.Price,
		Description: product.Description,
		Image:       product.Image,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpserver.Payload{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, httpserver.Payload{
		"product": respProduct,
	})
}

// DeleteProduct godoc
//
//	@Summary		Delete a product
//	@Description	remove the product information
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Product identifier"
//	@Success		200	{string}	string
//	@Failure		400	{string}	string	"when when bad request"
//	@Failure		500	{string}	string	"when delete product process error"
//	@Router			/product/{id}/ [delete]
func (pc *ProductController) DeleteProduct(c httpserver.HTTPContext) {
	id := c.Param("id")
	deleteProduct := products.NewDeleteProductUseCase(pc.productRepository)
	if err := deleteProduct.Execute(id); err != nil {
		c.JSON(http.StatusInternalServerError, httpserver.Payload{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, httpserver.Payload{
		"message": "Product deleted successfully",
	})
}

// ListProductsByCategory godoc
//
//	@Summary		List products
//	@Description	List a set of products information over followed categories
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			category	path		string	true	"Product category"
//	@Param			page		query		string	false	"Page defaults to 1"
//	@Param			size		query		string	false	"Size defaults to 50"
//	@Success		200			{array}		dtos.ProductDTO
//	@Failure		400			{string}	string	"when bad request"
//	@Failure		500			{string}	string	"when list products process error"
//	@Router			/product/{category}/ [get]
func (pc *ProductController) ListProductsByCategory(c httpserver.HTTPContext) {
	category := c.Param("category")

	if _, ok := errors.ValidCategories[category]; !ok {
		c.JSON(http.StatusBadRequest, httpserver.Payload{
			"error": errors.ErrInvalidCategory.Error(),
		})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "50"))

	products, err := products.NewListProductsByCategoryUseCase(pc.productRepository).Execute(category, page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpserver.Payload{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, httpserver.Payload{
		"products": products,
	})
}
