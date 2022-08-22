package main

import (
	"fashion-shop/controllers"
	"fashion-shop/initializers"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := httprouter.New()
	r.GET("/", controllers.GetAll)
	r.POST("/add-to-cart", controllers.AddToCart)

	log.Println("listen on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
