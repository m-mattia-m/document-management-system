package api

import (
	"document-management-system/pkg/document"
	"document-management-system/pkg/environment"
	"fmt"
	"github.com/swaggo/swag/example/basic/docs"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() {
	port, err := environment.GetEnv("APP_PORT")
	if err != nil {
		log.Println("[GET-ENV]: Get Port was failed \t\t -> Error: " + err.Error())
		return
	}

	r := gin.Default()
	host, err := environment.GetEnv("APP_HOST")
	if err == nil {
		docs.SwaggerInfo.Host = host
	}
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.Use(CORSMiddleware())
	v1 := r.Group("/api/v1", Auth)
	{
		documentsRouting := v1.Group("/documents")
		{
			documentsRouting.GET("/", document.UploadFile)
			documentsRouting.GET("/:id", document.UploadFile)
			documentsRouting.POST("/", document.UploadFile)
			documentsRouting.PUT("/:id", document.UploadFile)
			documentsRouting.DELETE("/:id", document.UploadFile)
		}
	}
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	err = r.Run(":" + port)
	if err != nil {
		portInt, err := strconv.ParseInt(port, 10, 64)
		if err != nil {
			log.Println("[SET-PORT]: Convert Port from String into Int was failed \t\t -> Error: " + err.Error())
			return
		}
		newPort := strconv.FormatUint(uint64(portInt+1), 10)
		log.Println("[SET-PORT]: Port '" + port + "' already in use \t\t.-> deploy on new port '" + newPort + "'")
		r.Run(":" + newPort)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Request-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

// TODO: Create Auth with Zitadel
func Auth(c *gin.Context) {
	user, password, hasAuth := c.Request.BasicAuth()
	_ = user
	_ = password
	_ = hasAuth

	fmt.Println("Login was successfully")
	// return true
}
