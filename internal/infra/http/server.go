package http

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

type Server struct {
	echo *echo.Echo
}

func NewServer() *Server {
	e := echo.New()
	return &Server{
		echo: e,
	}
}

func (s *Server) setupMiddlewares() {
	s.echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${remote_ip} | ${latency_human} | ${status} | ${method} ${path}\n\n",
	}))
	s.echo.Use(middleware.Recover())
	s.echo.Use(middleware.CORS())
	s.echo.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	s.echo.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(viper.GetString("JWT_SECRET")),
	}))
}

func (s *Server) StartServer() {
	s.setupMiddlewares()
	s.echo.Logger.Fatal(s.echo.Start(":" + viper.GetString("PORT")))
}
