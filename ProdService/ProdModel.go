package ProdService

import "strconv"

type ProdModel struct {
	ProdID   int    `json:"pid"`
	ProdName string `json:"pname"`
}

func NewProd(id int, prodName string) *ProdModel {
	return &ProdModel{ProdID: id, ProdName: prodName}
}
func NewProdList(n int) []*ProdModel {
	ret := make([]*ProdModel, 0)
	for i := 0; i < n; i++ {
		ret = append(ret, NewProd(100+i, "ProdName"+strconv.Itoa(100+i)))
	}
	return ret
}
