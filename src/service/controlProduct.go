package service

import (
	"github.com/crud/entities"
	"github.com/crud/lib"
	"github.com/crud/repository"
)

type ControlProduct struct {
	productDb *repository.ProductDb
}

func NewControlProduct() *ControlProduct {
	controlProduct := &ControlProduct{}
	controlProduct.Init()
	return controlProduct
}
func (cp *ControlProduct) Init() {

	cp.productDb = repository.NewProductDb()
	lib.Logger.Info("ControlProduct initialized!")
}
func (cp *ControlProduct) SaveProduct(prod *entities.ProductBody)   error {
	return cp.productDb.Save(prod)
}
