package product

import (
	"errors"
	"fmt"
	"time"

	"net/http"
	"path/filepath"
	"strconv"

	"github.com/Junx27/ho_go_day23/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ReadProductsHandler(c *gin.Context) {
	var products []model.Product

	query := `select * from products`
	filter := c.Query("filter")
	var args []any
	if filter != "" {
		query = fmt.Sprintf(
			"%s %s",
			query,
			"where name = ?",
		)

		args = append(args, filter)
	}

	err := model.DB.Raw(query, args...).Scan(&products).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Message: fmt.Sprintf("failed to save product: %s", err.Error()),
		})
		return
	}
	var productsDtos []model.ProductDto
	for _, product := range products {
		var productDto model.ProductDto
		productDto.FillFromModel(product)
		productsDtos = append(productsDtos, productDto)
	}

	c.JSON(http.StatusOK, model.Response{
		Success: true,
		Message: "Success",
		Data:    productsDtos,
	})
}

func ReadByIdProductsHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("invalid id: %s", idParam)),
		)
		return
	}

	var product model.Product
	err = model.DB.First(&product, id).Error
	statusCodeError := http.StatusInternalServerError
	if errors.Is(err, gorm.ErrRecordNotFound) {
		statusCodeError = http.StatusNotFound
	}

	if err != nil {
		c.JSON(statusCodeError, model.Response{
			Message: fmt.Sprintf("failed to get product: %s", err.Error()),
		})
		return
	}

	var productDto model.ProductDto
	productDto.FillFromModel(product)
	c.JSON(http.StatusOK, model.Response{
		Success: true,
		Message: "Success",
		Data:    productDto,
	})
}

func CreateProductHandler(c *gin.Context) {
	var productDto model.ProductDto
	err := c.ShouldBind(&productDto)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("failed to bind request: %s", err.Error())),
		)
		return
	}

	product := productDto.ToModel()
	err = model.DB.Create(&product).Error
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("failed to save product: %s", err.Error())),
		)
		return
	}

	productDto.ID = product.ID
	productDto.CreatedAt = product.CreatedAt
	productDto.UpdatedAt = product.UpdatedAt

	c.JSON(http.StatusOK, model.NewSuccessResponse("Success", productDto))
}

func UpdateProductHandler(c *gin.Context) {
	var productDto model.ProductDto
	err := c.ShouldBind(&productDto)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("failed to bind request: %s", err.Error())),
		)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("invalid id: %s", err.Error())),
		)
		return
	}

	var existingProduct model.Product
	err = model.DB.First(&existingProduct, id).Error
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			model.NewFailedResponse(fmt.Sprintf("product not found: %s", err.Error())),
		)
		return
	}

	product := productDto.ToModel()
	product.ID = existingProduct.ID
	product.CreatedAt = existingProduct.CreatedAt
	product.UpdatedAt = time.Now()
	err = model.DB.Save(&product).Error
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("failed to update product: %s", err.Error())),
		)
		return
	}
	productDto.ID = product.ID
	productDto.CreatedAt = product.CreatedAt
	productDto.UpdatedAt = product.UpdatedAt

	c.JSON(http.StatusOK, model.NewSuccessResponse("Success", productDto))
}

func DeleteProductHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("invalid id: %s", err.Error())),
		)
		return
	}

	product := model.Product{ID: id}
	result := model.DB.Delete(product)
	if result.Error != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("failed to delete product: %s", result.Error.Error())),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		model.NewSuccessResponse(fmt.Sprintf("%d products deleted", result.RowsAffected), nil),
	)
}

var productUploadDir = "uploads/products"

func UploadProductImageHandler(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("failed to get uploaded file: %s", err.Error())),
		)
		return
	}

	name := c.PostForm("name")
	path := filepath.Join(productUploadDir, name)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("failed to save file: %s", err.Error())),
		)
		return
	}

	var productDto model.ProductDto
	productDto.ImagePath = &path
	c.JSON(http.StatusOK, model.NewSuccessResponse("Success", productDto))
}
