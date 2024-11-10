package router

import (
	"github.com/Junx27/ho_go_day23/service/usecase/product"
	"github.com/Junx27/ho_go_day23/service/usecase/user"
	"github.com/gin-gonic/gin"
)

func ServerRoutes() {
	r := gin.Default()
	users := r.Group("/users/")
	{
		users.POST("/", user.CreateUserHandler)
		users.GET("/", user.ReadUsersHandler)
	}
	products := r.Group("/product/")
	{
		products.POST("/", product.CreateProductHandler)
		products.GET("/", product.ReadProductsHandler)
		products.GET("/:id", product.ReadByIdProductsHandler)
		products.PUT("/:id", product.UpdateProductHandler)
		products.DELETE("/:id", product.DeleteProductHandler)
		products.POST("/upload-product-image", product.UploadProductImageHandler)
	}
	_ = r.Run()
}
