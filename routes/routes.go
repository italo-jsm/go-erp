package routes

import (
	"github.com/italosm/go-erp/middlewares"
	"github.com/italosm/go-erp/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

//CreateRoutes create the routes
func CreateRoutes() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	productController := new(controllers.ProductController)
	router.Handle("/products", middlewares.ProductsMiddlewareHandler(http.HandlerFunc(productController.GetAllProducts))).Methods("GET")
	router.Handle("/products/cost", middlewares.ProductsMiddlewareHandler(http.HandlerFunc(productController.GetProductBySaleCost))).Queries("saleCost", "{saleCost}").Methods("GET")
	return router
}