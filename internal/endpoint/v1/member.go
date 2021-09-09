package endpointv1

import (
	"github.com/melvinvoetberg/ah-scanner-api/pkg/ah/member"
	"github.com/gin-gonic/gin"
)

func Member(c *gin.Context) {
	m := ahmember.GetMember()

	c.JSON(200, gin.H{
		"member": m,
	})
}
