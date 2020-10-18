package main

import (
	"sync"
	"github.com/italosm/go-erp/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	wg := &sync.WaitGroup{}
	ch := make(chan model.Product)
	wg.Add(1)
	go getProducts(wg, ch)
	fmt.Println("Started get products")
	for product := range ch{
		fmt.Println(product.Description)
	}
	wg.Wait()
}

func getProducts(wg *sync.WaitGroup, channel chan <- model.Product){
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
	wg.Done()
}