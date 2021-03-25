package controller

import (
	"code-example/src/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *RestController) AllProductCtrl(res *gin.Context) {
	var (
		produk []store.OutputProduk
	)
	model.GetAllProduct(&produk)

	res.JSON(200, produk)

}

func (r *RestController) AddProductCtrl(res *gin.Context) {
	var (
		produk store.TProduk
	)

	_ = res.BindJSON(&produk)

	model.AddProduct(&produk)
	switch produk.ID {
	case 0:
		res.JSON(http.StatusBadRequest, map[string]bool{
			"insert": false,
		})
	default:
		res.JSON(200, map[string]bool{
			"insert": true,
		})

	}

}

func (r *RestController) DetailProductCtrl(res *gin.Context) {
	var (
		id     = res.Param("product_id")
		produk store.OutputProduk
	)

	model.DetailProduct(&produk, id)
	switch produk.ID {
	case 0:
		res.JSON(http.StatusBadRequest, map[string]string{
			"error": "Product not Found",
		})
	default:
		res.JSON(200, produk)

	}
}

func (r *RestController) UpdateProductCtrl(res *gin.Context) {
	var (
		id     = res.Param("product_id")
		produk store.TProduk
		err    error
	)
	_ = res.BindJSON(&produk)
	model.UpdateProduct(&produk, &err, id)

	switch {
	case err != nil:
		res.JSON(http.StatusBadRequest, map[string]bool{
			"update": false,
		})
	default:
		res.JSON(200, map[string]bool{
			"update": true,
		})

	}
}

func (r *RestController) DeleteProductCtrl(res *gin.Context) {
	var (
		id  = res.Param("product_id")
		err error
	)

	model.DeleteProduct(&err, id)
	switch {
	case err != nil:
		res.JSON(http.StatusBadRequest, map[string]bool{
			"delete": false,
		})
	default:
		res.JSON(200, map[string]bool{
			"delete": true,
		})

	}
}
