package main

type (
	ProductModel struct {
		Id          int64  `json:"id"`
		Sku         string `json:"sku"`
		Description string `json:"description"`
	}
)

func NewProductModel(id int64, sku string, description string) ProductModel {
	return ProductModel{Id: id, Sku: sku, Description: description}
}
