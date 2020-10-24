package controllers

import (
	"github.com/italosm/go-erp/repository"
	"github.com/italosm/go-erp/model"
	"encoding/json"
	"net/http"
)

type ProductController struct{

}

func (productController *ProductController) GetAllProducts(w http.ResponseWriter, r *http.Request){
		products := []model.Product{}
		json.NewEncoder(w).Encode(products)
}

func (ProductController *ProductController) GetProductBySaleCost(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(repository.GetProductBySaleCost(5))
}
