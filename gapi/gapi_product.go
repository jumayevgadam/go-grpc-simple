package gapi

import (
	"context"
	"log"

	"github.com/jumayevgadam/go-grpc-simple/model"
	pb "github.com/jumayevgadam/go-grpc-simple/proto/product"
	"github.com/jumayevgadam/go-grpc-simple/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProductGapi struct {
	pb.UnimplementedProductGapiServer
	repository repository.ProductRepository
}

func NewGrpcProduct(repo repository.ProductRepository) *ProductGapi {
	return &ProductGapi{repository: repo}
}

func (s *ProductGapi) InsertProduct(ctx context.Context, input *pb.InsertProductRequest) (*pb.InsertProductResponse, error) {
	data := model.Product{
		Name:        input.Name,
		Price:       input.Price,
		IsAvailable: input.IsAvailable,
	}

	product, err := s.repository.InsertProduct(data)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	resp := pb.InsertProductResponse{
		Meta: &pb.Meta{
			Code:    200,
			Message: "Success",
		},
		Data: &pb.ProductDBResponse{
			Id:          product.Id,
			Name:        product.Name,
			Price:       product.Price,
			IsAvailable: product.IsAvailable,
			CreatedAt:   timestamppb.New(product.CreatedAt),
			UpdatedAt:   timestamppb.New(product.UpdatedAt),
		},
	}

	return &resp, nil
}

func (s *ProductGapi) GetListProducts(ctx context.Context, input *pb.GetListProductRequest) (*pb.GetListProductResponse, error) {
	if input.Page < 1 {
		input.Page = 1
	}

	if input.Limit < 1 {
		input.Limit = 10
	}

	filter := map[string]interface{}{
		"name": input.Name,
	}

	products, pagination, err := s.repository.GetListProducts(
		int(input.Limit), int(input.Page), input.Sorting, input.Direction, filter)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	resp := []*pb.ProductDBResponse{}
	for _, val := range products {
		product := pb.ProductDBResponse{
			Id:          val.Id,
			Name:        val.Name,
			Price:       val.Price,
			IsAvailable: val.IsAvailable,
			CreatedAt:   timestamppb.New(val.CreatedAt),
			UpdatedAt:   timestamppb.New(val.UpdatedAt),
		}

		resp = append(resp, &product)
	}

	result := pb.GetListProductResponse{
		Meta: &pb.Meta{
			Code:    200,
			Message: "Success",
			Kind: &pb.Meta_Pagination{
				Pagination: &pb.Pagination{
					Page:         int64(pagination.Page),
					Limit:        int64(pagination.Limit),
					TotalPage:    int64(pagination.TotalPage),
					TotalRecords: pagination.TotalRecords,
				},
			},
		},
		Data: resp,
	}

	return &result, nil
}

func (s *ProductGapi) GetProductByID(ctx context.Context, input *pb.GetProductByIDRequest) (*pb.GetProductByIDResponse, error) {
	if input.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "product id is empty")
	}

	product, err := s.repository.GetProductByID(input.Id)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result := pb.GetProductByIDResponse{
		Meta: &pb.Meta{
			Code:    200,
			Message: "Success",
		},
		Data: &pb.ProductDBResponse{
			Id:          product.Id,
			Name:        product.Name,
			Price:       product.Price,
			IsAvailable: product.IsAvailable,
			CreatedAt:   timestamppb.New(product.CreatedAt),
			UpdatedAt:   timestamppb.New(product.UpdatedAt),
		},
	}

	return &result, nil
}

func (s *ProductGapi) UpdateProductByID(ctx context.Context, input *pb.UpdateProductByIDRequest) (*pb.UpdateProductByIDResponse, error) {
	if input.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "product id is empty")
	}

	updateData := map[string]interface{}{
		"is_available": input.IsAvailable,
	}

	product, err := s.repository.UpdateProductByID(input.Id, updateData)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result := pb.UpdateProductByIDResponse{
		Meta: &pb.Meta{
			Code:    200,
			Message: "Success",
		},
		Data: &pb.ProductDBResponse{
			Id:          product.Id,
			Name:        product.Name,
			Price:       product.Price,
			IsAvailable: product.IsAvailable,
			CreatedAt:   timestamppb.New(product.CreatedAt),
			UpdatedAt:   timestamppb.New(product.UpdatedAt),
		},
	}

	return &result, nil
}

func (s *ProductGapi) DeleteProductByID(ctx context.Context, input *pb.DeleteProductByIDRequest) (*pb.DeleteProductByIDResponse, error) {
	if input.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "product id is empty")
	}

	err := s.repository.DeleteProductByID(input.Id)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result := pb.DeleteProductByIDResponse{
		Meta: &pb.Meta{
			Code:    200,
			Message: "Success",
		},
	}

	return &result, nil
}
