package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func CORSMiddleware() gin.HandlerFunc {
	//return func(c *gin.Context) {
	//	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://service.xamust.tech")
	//	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	//	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	//	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	//
	//	if c.Request.Method == "OPTIONS" {
	//		c.AbortWithStatus(http.StatusNoContent)
	//		return
	//	}
	//
	//	c.Next()
	//}
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://service.xamust.tech", "http://special.xamust.tech"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type, Authorization"},
		AllowCredentials: true,
		//AllowOriginFunc: func(origin string) bool {
		//	return origin == "https://github.com"
		//},
		MaxAge: 12 * time.Hour,
	})
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
