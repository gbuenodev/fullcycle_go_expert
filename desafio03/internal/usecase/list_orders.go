package usecase

import "github.com/gbuenodev/fullcycle_go_expert/desafio03/internal/entity"

type ListOrdersOutputDTO struct {
	Orders []OrderOutputDTO
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (lo *ListOrdersUseCase) Execute() (ListOrdersOutputDTO, error) {
	orders, err := lo.OrderRepository.GetAll()
	if err != nil {
		return ListOrdersOutputDTO{}, err
	}

	outputOrders := []OrderOutputDTO{}
	for _, order := range orders {
		orderDto := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		outputOrders = append(outputOrders, orderDto)
	}

	return ListOrdersOutputDTO{Orders: outputOrders}, nil
}
