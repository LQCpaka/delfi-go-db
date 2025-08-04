package handlers

import (
	"net/http"

	"delfi-scanner-api/db"
	"delfi-scanner-api/models"

	"github.com/gin-gonic/gin"
)

// AddProductToTicket add one product to exist ticket
func AddProductToTicket(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Validate if it's a valid ticket
	var ticket models.Ticket
	if result := db.DB.First(&ticket, product.ID); result != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Ticket ID"})
		return
	}

	if result := db.DB.Create(&product); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product added to ticket successfully", "product": product})
}

// UpdateProduct Cập nhật thông tin sản phẩm
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if result := db.DB.First(&product, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Bind dữ liệu mới, chỉ các trường được gửi sẽ được cập nhật
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	db.DB.Save(&product)
	c.JSON(http.StatusOK, product)
}

// DeleteProduct Xóa sản phẩm (xóa mềm)
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if result := db.DB.Delete(&models.Product{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
