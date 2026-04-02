package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mtk14m/manomano/internal/models"
)

// Mock liste produits
var productMockList = []models.Product{
	{
		ID:       1,
		Name:     "Perceuse visseuse",
		Price:    89.99,
		Category: "outillage",
		InStock:  true,
	},
	{
		ID:       2,
		Name:     "Aspirateur robot",
		Price:    249.99,
		Category: "électroménager",
		InStock:  false,
	},
	{
		ID:       3,
		Name:     "Casque audio sans fil",
		Price:    129.99,
		Category: "high-tech",
		InStock:  true,
	},
}

func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, productMockList)
}

func GetProductsById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid product id",
		})
		return
	}

	for _, product := range productMockList {
		if product.ID == id {
			//200
			c.JSON(http.StatusOK, product)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "product not found",
	})
}
