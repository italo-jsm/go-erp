package requests

import (
	"github.com/italosm/go-erp/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func RequestProducts() []model.Product{
	ch := make(chan model.Product)
	var products []model.Product
	go getProducts(ch)
	for product := range ch{
		products = append(products, product)
	}
	return products
}

func getProducts(channel chan <- model.Product){
	response, err := http.Get("http://localhost:8080/products")
	if(err != nil){
		panic(err.Error())
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	var products []model.Product
	json.Unmarshal(body, &products)
	for i := range products{
		channel <- products[i]
	}
	close(channel)
}