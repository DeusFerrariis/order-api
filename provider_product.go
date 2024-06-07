package main

func (provider *SqliteProvider) CreateProduct(sku string, description string) (*ProductModel, error) {
	res, err := provider.db.Exec("INSERT INTO products VALUES(NULL, ?, ?)", sku, description)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &ProductModel{
		Id:          id,
		Sku:         sku,
		Description: description,
	}, nil
}
