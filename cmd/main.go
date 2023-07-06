package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thifnmi/go-book-api/config"
	"github.com/thifnmi/go-book-api/pkg/interfaces/http/handlers"
	"github.com/thifnmi/go-book-api/pkg/interfaces/http/router/book"
	"github.com/thifnmi/go-book-api/pkg/interfaces/http/router/health"
	"github.com/thifnmi/go-book-api/pkg/repositories"
	"github.com/thifnmi/go-book-api/pkg/usecases"
	"github.com/thifnmi/go-book-api/utils"
)

func main() {
	var (
		mode int
	)
	conf := config.InitConfig()

	if strings.Contains(conf.Env, "development") {
		mode = 0
	} else if strings.Contains(conf.Env, "staging") {
		mode = 1
	} else {
		mode = 2
	}

	switch mode {
	case 0:
		gin.SetMode(gin.DebugMode)
	case 1:
		gin.SetMode(gin.TestMode)
	case 2:
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()

	// Config CORS for public API

	// corsConf := cors.New(cors.Options{
	// 	AllowedMethods: []string{"GET", "PUT", "PATCH", "OPTIONS", "POST", "DELETE", "HEAD", "DELETE"},
	// 	ExposedHeaders: []string{"Content-Length"},
	// 	AllowedOrigins: []string{
	// 		"http://*.bizflycloud.vn",
	// 		"http://*.vccloud.vn",
	// 		"https://*.bizflycloud.vn",
	// 		"https://*.vccloud.vn",
	// 		"http://*.bizflycloud.vn:8081",
	// 		"http://*.bizflycloud.vn:8080",
	// 		"http://local.bizflycloud.vn:8080",
	// 		"http://local.bizflycloud.vn:8000",
	// 		"http://hn-local.bizflycloud.vn",
	// 		"http://hcm-local.bizflycloud.vn",
	// 		"http://hn-local.bizflycloud.vn:8000",
	// 		"http://hn-local.bizflycloud.vn:8080",
	// 		"http://hcm-local.bizflycloud.vn:8000",
	// 		"http://hcm-local.bizflycloud.vn:8080",
	// 		"http://hn-local.manage.bizflycloud.vn",
	// 		"http://hcm-local.manage.bizflycloud.vn",
	// 		"http://hn-local.manage.bizflycloud.vn:8080",
	// 		"http://hcm-local.manage.bizflycloud.vn:8080",
	// 	},
	// 	AllowCredentials: true,
	// 	Debug:            true,
	// })

	// custom log format
	// r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	// 	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	// 		param.ClientIP,
	// 		param.TimeStamp.Format(time.RFC3339Nano),
	// 		param.Method,
	// 		param.Path,
	// 		param.Request.Proto,
	// 		param.StatusCode,
	// 		param.Latency,
	// 		param.Request.UserAgent(),
	// 		param.ErrorMessage,
	// 	)
	// }))
	db := utils.NewMysqlStorage(conf)

	r.Use(
		// corsConf, // enable if use config CORS
		gin.Logger(),
		gin.Recovery(),
	)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Page not found",
			"success": false,
		})
		c.Abort()
		return
	})
	r.HandleMethodNotAllowed = true

	bookRepo := repositories.NewbookRepository(db.Db())
	bookUseCase := usecases.NewBookUsecase(bookRepo)
	bookHandler := handlers.NewbookHandler(bookUseCase)

	healthhandler := handlers.NewHealthHandler()

	health.SetupRouter(r, *healthhandler)
	book.SetupRouter(r, *&bookHandler)

	if conf.Env == "production" {
		if err := r.Run(fmt.Sprintf("%s:%s", conf.ServerHost, conf.ServerPort)); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
		log.Printf("Server started at http://%s:%s\n", conf.ServerHost, conf.ServerPort)
	}
	if err := r.Run(fmt.Sprintf("%s:%s", conf.ServerHost, conf.ServerPort)); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
