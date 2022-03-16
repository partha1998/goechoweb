package tronics

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e = echo.New()

func init() {
	err := cleanenv.ReadEnv(&cfg)
	fmt.Printf("%+v", cfg)
	if err != nil {
		e.Logger.Fatal("uable to load configuration")
	}
}

func Start() {

	//custom middleware for all the router
	//e.Use(middlewareSample)

	//echo middleware
	e.Pre(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(1)))

	//middleware for only this router
	e.GET("/", sample, middlewareSample)

	e.GET("/product/:id", getProduct)

	e.GET("/products", getProducts)

	//echo middleware in specific router
	e.POST("/products", addProduct, middleware.BodyLimit("1k"))

	e.Logger.Print(fmt.Sprintf("The server listening on %v", cfg.Port))
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", cfg.Port)))
}
