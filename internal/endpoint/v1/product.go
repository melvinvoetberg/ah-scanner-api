package endpointv1

import (
	"github.com/melvinvoetberg/ah-scanner-api/pkg/ah/product"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ProductFIR(c *gin.Context) {
	fir, _ := strconv.ParseUint(c.Param("fir"), 10, 64)
	p := ahproduct.GetProduct(fir)

	c.JSON(200, gin.H{
		"product": p,
	})
}

func ProductGTIN(c *gin.Context) {
	gtin, _ := strconv.ParseUint(c.Param("gtin"), 10, 64)
	p := ahproduct.SearchGTIN(gtin)

	c.JSON(200, gin.H{
		"product": p,
	})
}
