package boot

import (
	"cscke/internal/route"
	"cscke/pkg/db"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	engine := gin.Default()

	route.Boot(engine)

	db.Boot()

	return engine
}
