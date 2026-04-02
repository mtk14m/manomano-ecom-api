package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	repositories "github.com/mtk14m/manomano/internal/repositories"
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
	c.JSON(http.StatusOK, h.repo.GetAll())
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid product id",
		})
		return
	}

	product, isProductFound := h.repo.GetByID(id)
	if !isProductFound {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "product not found",
		})
	} else {
		c.JSON(http.StatusOK, product)
	}
}
