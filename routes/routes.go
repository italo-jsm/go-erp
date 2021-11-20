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
	loginController := new(controllers.LoginController)
	router.HandleFunc("/login", func (w http.ResponseWriter, r *http.Request){
		loginController.AttemptAuthentication(w, r)
	}).Methods("POST")
	router.Handle("/products", middlewares.ProductsMiddlewareHandler(http.HandlerFunc(productController.GetAllProducts))).Methods("GET")
	router.Handle("/products/cost", middlewares.ProductsMiddlewareHandler(http.HandlerFunc(productController.GetProductBySaleCost))).Queries("saleCost", "{saleCost}").Methods("GET")
	return router
}