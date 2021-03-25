package main

import (
	"code-example/src/middleware"
	"code-example/src/routers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
func main() {
	mode := os.Getenv("MODE")
	gin.SetMode(mode)

	var (
		// setup routes
		route  = gin.Default()
		auth   = route.Group("/auth")
		produk = route.Group("/products")

		// initial package used
		app    routers.RestRouter
		middle middleware.RestMiddleware
	)

	//used middleware
	produk.Use(
		middle.ReqValidator,
	)

	// used routes
	app.IndexRoute(route)
	app.AuthRoute(auth)
	app.ProdukRoute(produk)

	_ = route.Run()
}
