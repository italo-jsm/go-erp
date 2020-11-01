package controllers

import (
	"strconv"
	"github.com/gorilla/mux"
	"github.com/italosm/go-erp/repository"
	"encoding/json"
	"net/http"
)

type ProductController struct{

}

func (productController *ProductController) GetAllProducts(w http.ResponseWriter, r *http.Request){
	products := repository.GetAllProducts()
	json.NewEncoder(w).Encode(products)
}

func (ProductController *ProductController) GetProductBySaleCost(w http.ResponseWriter, r *http.Request){
	productSaleCost := mux.Vars(r)["saleCost"]
	numberSaleCost, _ := strconv.Atoi(productSaleCost)
	repository.GetProductBySaleCost(numberSaleCost)
	json.NewEncoder(w).Encode(repository.GetProductBySaleCost(5))
}
