package auth

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/xamops/auth/pkg/auth"
	"log"
	"net/http"
)

const (
	accessToken  = "access_token"
	refreshToken = "refresh_token"
)

func Auth(token auth.TokenHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		access, err := c.Cookie(accessToken)
		if err != nil {
			log.Print("access err:", err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		if _, err := c.Cookie(refreshToken); err != nil {
			log.Print("refresh err:", err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		parse, err := token.Parse(access)
		if err != nil {
			log.Print("parse err:", err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		ctx := auth.Context(c.Request.Context(), parse)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func Refresh(token auth.TokenHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		refresh, err := c.Cookie(refreshToken)
		if err != nil {
			log.Print("access err:", err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		parse, err := token.Parse(refresh)
		if err != nil {
			log.Print("parse err:", err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		ctx := auth.Context(c.Request.Context(), parse)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
