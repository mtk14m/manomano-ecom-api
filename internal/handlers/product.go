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
	"github.com/mtk14m/manomano/internal/repositories"
)

type ProductHandler struct {
	repo *repositories.ProductRepository
}

func NewProductHandler(repo *repositories.ProductRepository) *ProductHandler {
	return &ProductHandler{
		repo: repo,
	}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid product id",
		})
		return
	}

	product, err := h.repo.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "product not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return

	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var dto dtos.CreateProductDto

	if err := c.ShouldBindJSON(&dto); err != nil {
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
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request payload",
		})
		return
	}

	p := models.Product{
		Name:     dto.Name,
		Price:    dto.Price,
		Category: dto.Category,
		InStock:  dto.InStock,
	}

	product, err := h.repo.Create(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid product id",
		})
		return
	}
	var dto dtos.UpdateProductDto

	if err := c.ShouldBindJSON(&dto); err != nil {
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
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request payload",
		})
		return
	}

	p := models.Product{
		Name:     dto.Name,
		Price:    dto.Price,
		Category: dto.Category,
		InStock:  dto.InStock,
	}

	product, err := h.repo.Update(id, p)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "product not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, product)
}
