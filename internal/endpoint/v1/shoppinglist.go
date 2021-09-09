package endpointv1

import (
	"github.com/melvinvoetberg/ah-scanner-api/pkg/ah/shoppinglist"
	"github.com/gin-gonic/gin"
)

type ShoppinglistAddBody struct {
	Type string `json:"type"`
	Id int `json:"id"`
}

func Shoppinglist(c *gin.Context) {
	s := ahshoppinglist.GetShoppinglist()

	c.JSON(200, gin.H{
		"shoppinglist": s,
	})
}

func ShoppinglistAdd(c *gin.Context) {
	var b ShoppinglistAddBody
	var s ahshoppinglist.Shoppinglist

	if c.ShouldBind(&b) == nil {
		if b.Type == "FIR" {
			s = ahshoppinglist.AddProduct(b.Id)
		} else if b.Type == "GTIN" {
			s = ahshoppinglist.AddProductGTIN(b.Id)
		}

		c.JSON(200, gin.H{
			"shoppinglist": s,
		})
  }
}
