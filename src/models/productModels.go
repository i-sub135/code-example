package models

import "code-example/src/store"

func (r *RestModels) AddProduct(data *store.TProduk) {
	sql.Create(data)
}
func (r *RestModels) GetAllProduct(data *[]store.OutputProduk) {
	sql.Table("t_produk").
		Order("rating DESC").
		Find(data)
}

func (r *RestModels) DetailProduct(data *store.OutputProduk, id string) {
	sql.Table("t_produk").
		Where("id = ?", id).
		Find(data)
}

func (r *RestModels) UpdateProduct(data *store.TProduk, err *error, id string) {
	*err = sql.Table("t_produk").
		Where("id = ?", id).
		Updates(data).Error
}

func (r *RestModels) DeleteProduct(err *error, id string) {
	*err = sql.Table("t_produk").
		Where("id = ?", id).Delete(map[string]string{}).
		Error
}
