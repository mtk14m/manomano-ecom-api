package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mtk14m/manomano/internal/dtos"
	"github.com/mtk14m/manomano/internal/models"
	"github.com/mtk14m/manomano/internal/services"
)

type ProductHandler struct {
	productService *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: service,
	}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.productService.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id, err := parseProductID(c)
	if err != nil {
		return
	}

	product, err := h.productService.GetProductByID(id)
	if handleRepoError(c, err) {
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var dto dtos.CreateProductDto

	if !bindAndValidateJSON(c, &dto) {
		return
	}

	p := models.Product{
		Name:     dto.Name,
		Price:    dto.Price,
		Category: dto.Category,
		InStock:  dto.InStock,
	}

	product, err := h.productService.CreateProduct(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {

	var dto dtos.UpdateProductDto
	id, err := parseProductID(c)
	if err != nil {
		return
	}

	if !bindAndValidateJSON(c, &dto) {
		return
	}
	p := models.Product{
		Name:     dto.Name,
		Price:    dto.Price,
		Category: dto.Category,
		InStock:  dto.InStock,
	}

	product, err := h.productService.UpdateProduct(id, p)
	if handleRepoError(c, err) {
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := parseProductID(c) // on valide le param
	if err != nil {
		return
	}
	err = h.productService.DeleteProduct(id)
	if handleRepoError(c, err) {
		return
	}

	c.Status(http.StatusNoContent)
}

//HELPERS

func bindAndValidateJSON(c *gin.Context, dto any) bool {

	if err := c.ShouldBindJSON(dto); err != nil {
		var validationErrs validator.ValidationErrors

		if errors.As(err, &validationErrs) {
			out := make(map[string]string)

			for _, fieldErr := range validationErrs {
				field := strings.ToLower(fieldErr.Field())

				switch fieldErr.Tag() {
				case "required":
					out[field] = "this field is required"
				case "min":
					out[field] = fmt.Sprintf("must be at least %s", fieldErr.Param())
				case "max":
					out[field] = fmt.Sprintf("must be at most %s", fieldErr.Param())
				case "gte":
					out[field] = fmt.Sprintf("must be greater than or equal to %s", fieldErr.Param())
				case "gt":
					out[field] = fmt.Sprintf("must be greater than %s", fieldErr.Param())
				case "lte":
					out[field] = fmt.Sprintf("must be less than or equal to %s", fieldErr.Param())
				default:
					out[field] = fmt.Sprintf("failed on '%s' validation", fieldErr.Tag())
				}
			}

			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "validation failed",
				"details": out,
			})
			return false
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request payload",
		})
		return false
	}

	return true
}

func parseProductID(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid product id",
		})
		return -1, err
	}
	return id, nil
}

func handleRepoError(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "product not found",
		})
		return true
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "internal server error",
	})
	return true
}
