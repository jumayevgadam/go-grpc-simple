package repository

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/jumayevgadam/go-grpc-simple/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// var _ ProductRepository = (*ProductRepositoryImpl)(nil)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

type ProductRepository interface {
	InsertProduct(product model.Product) (*model.Product, error)
	GetListProducts(limit, page int, sorting, direction string, filter map[string]interface{}) (
		[]model.Product, *model.Pagination, error,
	)
	GetProductByID(id string) (*model.Product, error)
	UpdateProductByID(id string, input map[string]interface{}) (*model.Product, error)
	DeleteProductByID(id string) error
}

func NewDataProduct(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db}
}

func (r *ProductRepositoryImpl) InsertProduct(product model.Product) (*model.Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return nil, fmt.Errorf("[repository][InsertProduct]: %v", err.Error())
	}

	return &product, nil
}

func (r *ProductRepositoryImpl) GetListProducts(limit, page int, sorting, direction string, filter map[string]interface{}) (
	[]model.Product, *model.Pagination, error,
) {
	var products []model.Product
	var pagination model.Pagination

	query := r.db

	if filter["name"] != "" {
		query = query.Where("name = ?", filter["name"].(string))
	}

	pagination.Limit = limit
	pagination.Page = page

	if strings.EqualFold(direction, "asc") {
		direction = "asc"
	}

	query = query.Order(fmt.Sprintf("%s %s", sorting, direction))

	err := query.Scopes(r.Paginate(products, &pagination, query, int64(len(products)))).Find(&products).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, nil
		}

		return nil, nil, fmt.Errorf("[query][Scopes]: %v", err.Error())
	}

	return products, &pagination, nil
}

func (r *ProductRepositoryImpl) GetProductByID(id string) (*model.Product, error) {
	var product model.Product
	err := r.db.Model(&model.Product{}).Where("id = ?", id).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err)
			return nil, err
		}

		log.Println(err)
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepositoryImpl) UpdateProductByID(id string, input map[string]interface{}) (*model.Product, error) {
	var product model.Product
	err := r.db.Model(&product).Where("id = ?", id).Clauses(clause.Returning{}).Updates(input).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepositoryImpl) DeleteProductByID(id string) error {
	err := r.db.Where("id = ?", id).Delete(&model.Product{}).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *ProductRepositoryImpl) Paginate(value interface{}, pagination *model.Pagination, db *gorm.DB,
	currRecord int64) func(db *gorm.DB) *gorm.DB {
	var totalRecords int64
	db.Model(value).Count(&totalRecords)

	pagination.TotalRecords = totalRecords
	pagination.TotalPage = int(math.Ceil(float64(totalRecords) / float64(pagination.GetLimit())))
	pagination.Records = int64(pagination.Limit*(pagination.Page-1)) + int64(currRecord)

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit())
	}
}
