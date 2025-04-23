package main

import (
    "github.com/gin-gonic/gin"
    "product-api/config"
    "product-api/router"
)

func main() {
    db := config.InitDB()
    r := gin.Default()

    router.InitializeRoutes(r, db)

    r.Run(":8080") // Inicia o servidor na porta 8080
}

