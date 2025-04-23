package controller

import (
    "net/http"
    "product-api/service"
    "github.com/gin-gonic/gin"
)

type ProductController struct {
    Service *service.ProductService
}

func (pc *ProductController) GetAll(c *gin.Context) {
    products, err := pc.Service.GetAllProducts()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar produtos"})
        return
    }
    c.JSON(http.StatusOK, products)
}

