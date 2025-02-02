package handlers

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/xamops/auth/internal/dto"
	"gitlab.com/xamops/auth/pkg/auth"
	"log/slog"
	"net/http"
)

// @BasePath /api/v1

// Login godoc
// @Summary Login
// @Schemes
// @Description Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "query params"
// @Success 200 {object} dto.AuthResponse
// @Router /login [post]
func (a *apiHandlersV1) Login(c *gin.Context) {
	ctx := c.Request.Context()
	req := dto.LoginRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	user, err := a.usecases.Auth().Login(ctx, req.Email, req.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
	a.logger.Info("login success", "user:", user.Email)

	claims := a.mappers.Users().ToClaims(user)
	access, err := a.access.Generate(claims)
	if err != nil {
		a.logger.Error("failed generate access token", slog.String("error", err.Error()))
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err})
	}

	refresh, err := a.refresh.Generate(a.mappers.Users().ToClaims(user))
	if err != nil {
		a.logger.Error("failed generate refresh token", slog.String("error", err.Error()))
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err})
	}

	c.SetCookie(accessToken, access, int(claims.ExpiresAt), "", "", false, false)
	c.SetCookie(refreshToken, refresh, int(claims.ExpiresAt), "", "", false, false)

	c.SecureJSON(http.StatusOK, dto.AuthResponse{
		ID:  a.mappers.UUID().ToString(claims.UUID),
		TTL: claims.ExpiresAt,
	})
}

// Logout godoc
// @Summary Logout
// @Schemes
// @Description Logout
// @Tags Auth
// @Accept json
// @Produce json
// @Success 204
// @Router /logout [post]
func (a *apiHandlersV1) Logout(c *gin.Context) {
	c.SetCookie(accessToken, "", 0, "", "", false, false)
	c.SetCookie(refreshToken, "", 0, "", "", false, false)
	c.JSON(http.StatusNoContent, nil)
}

// Refresh godoc
// @Summary Refresh
// @Schemes
// @Description Refresh
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} dto.AuthResponse
// @Router /refresh [post]
func (a *apiHandlersV1) Refresh(c *gin.Context) {
	ctx := c.Request.Context()
	claims := auth.FromContext(ctx)
	access, err := a.access.Generate(claims)
	if err != nil {
		a.logger.Error("failed generate refresh token", slog.String("error", err.Error()))
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err})
	}
	c.SetCookie(accessToken, access, int(claims.ExpiresAt), "", "", false, false)
	c.SecureJSON(http.StatusOK, dto.AuthResponse{
		ID:  a.mappers.UUID().ToString(claims.UUID),
		TTL: claims.ExpiresAt,
	})
}
