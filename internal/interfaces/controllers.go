package interfaces

import "github.com/gin-gonic/gin"

type Handlers interface {
	Login(*gin.Context)
	Logout(*gin.Context)
	Refresh(*gin.Context)

	GetCurrentUser(c *gin.Context)
}
