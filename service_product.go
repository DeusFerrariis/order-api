package main

type ProductService interface {
	CreateProduct(sku string, description string) (ProductModel, error)
}
