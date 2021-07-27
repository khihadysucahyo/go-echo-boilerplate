package routes

import (
	sentryecho "github.com/getsentry/sentry-go/echo"

	"github.com/khihadysucahyo/go-echo-boilerplate/repositories"
	s "github.com/khihadysucahyo/go-echo-boilerplate/server"
	"github.com/khihadysucahyo/go-echo-boilerplate/server/handlers"

	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func ConfigureRoutes(server *s.Server) {
	postHandler := handlers.NewPostHandlers(server)
	authHandler := handlers.NewAuthHandler(server)
	registerHandler := handlers.NewRegisterHandler(server)

	sentryMiddleware := sentryecho.New(sentryecho.Options{
		Repanic: true,
	})

	server.Echo.Use(middleware.Logger())
	server.Echo.Use(middleware.Recover())
	server.Echo.Use(sentryMiddleware)

	server.Echo.GET("/swagger/*", echoSwagger.WrapHandler)

	server.Echo.POST("/login", authHandler.Login)
	server.Echo.POST("/register", registerHandler.Register)
	server.Echo.POST("/refresh", authHandler.RefreshToken)

	r := server.Echo.Group("")
	config := middleware.JWTConfig{
		Claims:     &repositories.JwtCustomClaims{},
		SigningKey: []byte(server.Config.Auth.AccessSecret),
	}
	r.Use(middleware.JWTWithConfig(config))

	r.GET("/posts", postHandler.GetPosts)
	r.POST("/posts", postHandler.CreatePost)
	r.DELETE("/posts/:id", postHandler.DeletePost)
	r.PUT("/posts/:id", postHandler.UpdatePost)
}
