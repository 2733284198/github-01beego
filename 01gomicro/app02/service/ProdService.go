package service

import "strconv"

type ProdModel struct {
	ProdID   int
	ProdName string
}

func NewProd(id int, pname string) *ProdModel {
	return &ProdModel{
		ProdID:   id,
		ProdName: pname,
	}
}

func NewProdList(n int) []*ProdModel {
	ret := make([]*ProdModel, 0)

	for i := 0; i < n; i++ {
		ret = append(ret, NewProd(100+i, "prodName"+strconv.Itoa(100+i)))
	}

	return ret
}
