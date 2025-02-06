package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//c.Writer.Header().Set("Access-Control-Allow-Origin", "https://yourfrontend.com")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// to NGINX
//
//proxy_set_header Host $host;
//        proxy_set_header X-Real-IP $remote_addr;
//        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
//
//        # CORS headers
//        add_header Access-Control-Allow-Origin "*" always;
//        add_header Access-Control-Allow-Methods "GET, POST, PUT, DELETE, OPTIONS" always;
//        add_header Access-Control-Allow-Headers "Content-Type, Authorization" always;
//
//        # Обработка OPTIONS-запросов
//        if ($request_method = OPTIONS) {
//            return 204;
//        }
