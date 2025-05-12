package project

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Product struct {
	Id          int     `json:"id"`
	ProductName string  `json:"productName"`
	CategoryId  int     `json:"categoryId"`
	UnitPrice   float64 `json:"price"`
}

type Category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

func GetAllProducts() {
	response, err := http.Get("http://localhost:3000/products")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	bodyBytes, _ := io.ReadAll(response.Body)

	var products []Product
	json.Unmarshal(bodyBytes, &products)
	fmt.Println(products)

}

func AddProduct() {
	product := Product{
		// Id:          7,
		ProductName: "mikrofon",
		CategoryId:  2,
		UnitPrice:   100.0,
	}

	jsonProduct, err := json.Marshal(product)

	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	bodyBytes, _ := io.ReadAll(response.Body)

	var productResponse Product

	json.Unmarshal(bodyBytes, &productResponse)

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Kaydedildi: ", productResponse)
}
