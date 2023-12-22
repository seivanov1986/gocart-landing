package repository

import (
	"github.com/seivanov1986/sql_client"

	"github.com/seivanov1986/gocart/internal/repository/attribute"
	"github.com/seivanov1986/gocart/internal/repository/attribute_to_product"
	"github.com/seivanov1986/gocart/internal/repository/category"
	"github.com/seivanov1986/gocart/internal/repository/category_closure"
	"github.com/seivanov1986/gocart/internal/repository/image_to_category"
	"github.com/seivanov1986/gocart/internal/repository/image_to_product"
	"github.com/seivanov1986/gocart/internal/repository/meta"
	"github.com/seivanov1986/gocart/internal/repository/page"
	"github.com/seivanov1986/gocart/internal/repository/product"
	"github.com/seivanov1986/gocart/internal/repository/product_to_category"
	"github.com/seivanov1986/gocart/internal/repository/sefurl"
	"github.com/seivanov1986/gocart/internal/repository/user"
)

type Hub interface {
	Attribute() attribute.Repository
	AttributeToProduct() attribute_to_product.Repository
	Category() category.Repository
	CategoryClosure() category_closure.Repository
	ImageToCategory() image_to_category.Repository
	ImageToProduct() image_to_product.Repository
	Meta() meta.Repository
	Page() page.Repository
	Product() product.Repository
	ProductToCategory() product_to_category.Repository
	SefUrl() sefurl.Repository
	User() user.Repository
}

type hub struct {
	attribute          attribute.Repository
	attributeToProduct attribute_to_product.Repository
	category           category.Repository
	categoryClosure    category_closure.Repository
	imageToCategory    image_to_category.Repository
	imageToProduct     image_to_product.Repository
	meta               meta.Repository
	page               page.Repository
	product            product.Repository
	productToCategory  product_to_category.Repository
	sefUrl             sefurl.Repository
	user               user.Repository
}

func New(db sql_client.DataBase, trx sql_client.TransactionManager) *hub {
	return &hub{
		attribute:          attribute.New(db),
		attributeToProduct: attribute_to_product.New(db),
		category:           category.New(db, trx),
		categoryClosure:    category_closure.New(db, trx),
		imageToCategory:    image_to_category.New(db),
		imageToProduct:     image_to_product.New(db),
		meta:               meta.New(db, trx),
		product:            product.New(db, trx),
		productToCategory:  product_to_category.New(db),
		sefUrl:             sefurl.New(db, trx),
		user:               user.New(db),
	}
}

func (h *hub) Attribute() attribute.Repository {
	return h.attribute
}

func (h *hub) AttributeToProduct() attribute_to_product.Repository {
	return h.attributeToProduct
}

func (h *hub) Category() category.Repository {
	return h.category
}

func (h *hub) CategoryClosure() category_closure.Repository {
	return h.categoryClosure
}

func (h *hub) ImageToCategory() image_to_category.Repository {
	return h.imageToCategory
}

func (h *hub) ImageToProduct() image_to_product.Repository {
	return h.imageToProduct
}

func (h *hub) Meta() meta.Repository {
	return h.meta
}

func (h *hub) SefUrl() sefurl.Repository {
	return h.sefUrl
}

func (h *hub) Page() page.Repository {
	return h.page
}

func (h *hub) Product() product.Repository {
	return h.product
}

func (h *hub) ProductToCategory() product_to_category.Repository {
	return h.productToCategory
}

func (h *hub) User() user.Repository {
	return h.user
}
