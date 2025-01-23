package handlers

import (
	"github.com/gin-gonic/gin"
	_ "gitlab.com/xamops/auth/internal/dto"
	"gitlab.com/xamops/auth/pkg/auth"
	"net/http"
)

// @BasePath /api/v1

// GetCurrentUser godoc
// @Summary Get Current User
// @Schemes
// @Description Get Current User
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @securityDefinitions.basic BasicAuth
// @in header
// @Success 200 {object} dto.User
// @Router /user/current [get]
func (a *apiHandlersV1) GetCurrentUser(c *gin.Context) {
	ctx := c.Request.Context()
	claims := auth.FromContext(ctx)
	usr, err := a.usecases.Users().GetByUUID(ctx, claims.UUID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	user, err := a.mappers.Users().ToDTO(usr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, user)
}
