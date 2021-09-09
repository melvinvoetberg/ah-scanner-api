package main

import (
	"github.com/melvinvoetberg/ah-scanner-api/internal/endpoint/v1"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/product/:fir", endpointv1.ProductFIR)
		v1.GET("/product/gtin/:gtin", endpointv1.ProductGTIN)
		v1.GET("/member", endpointv1.Member)
		v1.GET("/shoppinglist", endpointv1.Shoppinglist)
		v1.POST("/shoppinglist/add", endpointv1.ShoppinglistAdd)
	}

	router.Run()
}
