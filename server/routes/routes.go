package routes

import (
	"fmt"

	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/khihadysucahyo/go-echo-boilerplate/services/token"

	s "github.com/khihadysucahyo/go-echo-boilerplate/server"
	"github.com/khihadysucahyo/go-echo-boilerplate/server/handlers"

	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	sentryTransaction "github.com/khihadysucahyo/go-echo-boilerplate/middleware"
)

func ConfigureRoutes(server *s.Server) {
	postHandler := handlers.NewPostHandlers(server)
	authHandler := handlers.NewAuthHandler(server)
	registerHandler := handlers.NewRegisterHandler(server)

	// record transaction for sentry apm
	sentryMiddleware := sentryTransaction.InitMiddleware()
	server.Echo.Use(sentryMiddleware.CORS)
	server.Echo.Use(sentryMiddleware.SENTRY)

	server.Echo.Use(middleware.Logger())

	server.Echo.Use(sentryecho.New(sentryecho.Options{
		Repanic: true,
	}))

	server.Echo.GET("/swagger/*", echoSwagger.WrapHandler)

	server.Echo.POST("/login", authHandler.Login)
	server.Echo.POST("/register", registerHandler.Register)
	server.Echo.POST("/refresh", authHandler.RefreshToken)

	fmt.Println(server.Config.Auth.AccessSecret)

	r := server.Echo.Group("")
	config := middleware.JWTConfig{
		Claims:     &token.JwtCustomClaims{},
		SigningKey: []byte(server.Config.Auth.AccessSecret),
	}
	r.Use(middleware.JWTWithConfig(config))

	r.GET("/posts", postHandler.GetPosts)
	r.POST("/posts", postHandler.CreatePost)
	r.DELETE("/posts/:id", postHandler.DeletePost)
	r.PUT("/posts/:id", postHandler.UpdatePost)
}
