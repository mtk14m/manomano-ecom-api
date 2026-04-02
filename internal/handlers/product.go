package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mtk14m/manomano/internal/models"
)

func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, models.ProductMockList)
}

func GetProductsById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid product id",
		})
		return
	}

	for _, product := range models.ProductMockList {
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
