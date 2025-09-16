package services

import (
	"context"
	"grpc/internal/pb"
	"grpc/internal/store"

	"io"

	"google.golang.org/grpc"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB store.Category
}

func NewCategoryService(categoryDB store.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	categoriesResponse := make([]*pb.Category, 0)

	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	return &pb.CategoryList{Categories: categoriesResponse}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.GetCategoryByIDRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.FindByID(in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (c *CategoryService) CreateCategoryStream(stream grpc.ClientStreamingServer[pb.CreateCategoryRequest, pb.CategoryList]) error {
	var categories []*pb.Category
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&pb.CategoryList{Categories: categories})
			}
			return err
		}
		category, err := c.CategoryDB.Create(req.Name, req.Description)
		if err != nil {
			return err
		}
		categories = append(categories, &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}
}

func (c *CategoryService) CreateCategoryStreamBidirectional(stream grpc.BidiStreamingServer[pb.CreateCategoryRequest, pb.Category]) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		category, err := c.CategoryDB.Create(req.Name, req.Description)
		if err != nil {
			return err
		}
		resp := &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
}
