package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mtk14m/manomano/internal/handlers"
	"github.com/mtk14m/manomano/internal/repositories"
	"github.com/mtk14m/manomano/internal/services"
	"github.com/mtk14m/manomano/pkg/database"
)

func main() {
	//init de la db
	db, err := database.NewDB()
	if err != nil {
		log.Fatalf("ERROR-DB: %s", err.Error())
	}
	defer db.Close()

	//On init le router et handlers de base de l'apps
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Ok, cool raoul",
		})
	})
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	//Création et injection des dépendances
	productRepository := repositories.NewProductRepository(db)
	productService:= services.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService) 
	r.GET("/products", productHandler.GetProducts)
	r.GET("/products/:id", productHandler.GetProductByID)
	r.POST("/products", productHandler.CreateProduct)
	r.PUT("/products/:id", productHandler.UpdateProduct)
	r.DELETE("/products/:id", productHandler.DeleteProduct)

	//on lance le server et on check si erreur
	log.Printf("App is running on port: 8000")
	if err = r.Run(":8000"); err != nil {
		log.Fatalf("ERROR-HTTP: %v", err)
	}

}
