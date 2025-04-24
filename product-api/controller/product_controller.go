package controller

import (
	"net/http"
	"product-api/model"
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

func (pc *ProductController) GetByID(c *gin.Context) {
	id := c.Param("id")
	product, err := pc.Service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (pc *ProductController) GetByName(c *gin.Context) {
	name := c.Query("name")
	products, err := pc.Service.GetByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar produtos"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) Count(c *gin.Context) {
	count, err := pc.Service.CountProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao contar produtos"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}

func (pc *ProductController) Create(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}
	if err := pc.Service.CreateProduct(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar produto"})
		return
	}
	c.JSON(http.StatusCreated, product)
}

func (pc *ProductController) Update(c *gin.Context) {
	id := c.Param("id")
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}
	if err := pc.Service.UpdateProduct(id, product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar produto"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (pc *ProductController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := pc.Service.DeleteProduct(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar produto"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
