package cheers

import (
	"github.com/gin-gonic/gin"
	"github.com/ginmills/ginmill"
)

type ICheers interface {

	// GET /ginmills/cheers
	Cheers(c *gin.Context)
}

// Cheers features
func Features(gm ICheers) (features *ginmill.Features) {
	r := gin.New()

	r.GET("/ginmills/cheers", gm.Cheers)
	features = ginmill.NewFeatures(r.Routes())

	return features
}
