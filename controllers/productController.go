package controllers

import (
	"encoding/json"
	"github.com/italosm/go-erp/requests"
	"net/http"
)

type ProductController struct{

}

func (productController *ProductController) GetAllProducts(w http.ResponseWriter, r *http.Request){
		products := requests.RequestProducts()
		json.NewEncoder(w).Encode(products)
}