package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	clog "github.com/charmbracelet/log"
)

type ProductController struct {
	service ProductService
}

func (pc *ProductController) CreateProduct(rw http.ResponseWriter, r *http.Request) {
	var body struct {
		Sku         string `json:"sku"`
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		rw.WriteHeader(400)
		fmt.Sprintln(rw, err.Error())
		return
	}
	product, err := pc.service.CreateProduct(body.Sku, body.Description)
	if err != nil {
		clog.Error(err)
		return
	}

	json.NewEncoder(rw).Encode(product)
}
