package order

import (
	"fmt"
	_entities "group-project/limamart/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllOrder(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		orderUseCase := NewOrderUseCase(mockOrderRepository{}, mockCartRepository{}, mockProductRepository{})
		data, rows, err := orderUseCase.GetAllOrder(1)
		assert.Nil(t, nil, err)
		assert.Equal(t, uint(1), data[0].UserID)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		orderUseCase := NewOrderUseCase(mockOrderRepositoryError{}, mockCartRepositoryError{}, mockProductRepositoryError{})
		data, rows, err := orderUseCase.GetAllOrder(1)
		assert.NotNil(t, err)
		assert.Nil(t, data)
		assert.Nil(t, nil, rows)
	})
}

func TestCrateOrder(t *testing.T) {
	t.Run("TestCreateOrderSuccess", func(t *testing.T) {
		orderUseCase := NewOrderUseCase(mockOrderRepository{}, mockCartRepository{}, mockProductRepository{})
		data, rows, err := orderUseCase.CreateOrder(_entities.Order{}, []uint{1, 2}, 2)
		assert.Nil(t, nil, err)
		assert.Equal(t, uint(1), data.UserID)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestCreateOrderError", func(t *testing.T) {
		orderUseCase := NewOrderUseCase(mockOrderRepositoryError{}, mockCartRepositoryError{}, mockProductRepositoryError{})
		data, rows, err := orderUseCase.CreateOrder(_entities.Order{}, []uint{1, 2}, 2)
		assert.NotNil(t, err)
		assert.Nil(t, nil, data)
		assert.Nil(t, nil, rows)
	})
}

func TestGetHistoryByID(t *testing.T) {
	t.Run("TestGetHistorySuccess", func(t *testing.T) {
		orderUseCase := NewOrderUseCase(mockOrderRepository{}, mockCartRepository{}, mockProductRepository{})
		data, rows, err := orderUseCase.GetHistoriOrderbyID(2)
		assert.Nil(t, nil, err)
		assert.Equal(t, uint(1), data.UserID)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestGetHistoryError", func(t *testing.T) {
		orderUseCase := NewOrderUseCase(mockOrderRepositoryError{}, mockCartRepositoryError{}, mockProductRepositoryError{})
		data, rows, err := orderUseCase.GetHistoriOrderbyID(2)
		assert.NotNil(t, err)
		assert.Nil(t, nil, data)
		assert.Nil(t, nil, rows)
	})
}

func TestCancelOrder(t *testing.T) {
	t.Run("TestcancelOrderSuccess", func(t *testing.T) {
		orderUseCase := NewOrderUseCase(mockOrderRepository{}, mockCartRepository{}, mockProductRepository{})
		data, rows, err := orderUseCase.CancelOrder(_entities.OrdersDetail{UserID: 1}, 1, 1)
		assert.Nil(t, nil, err)
		assert.NotEqual(t, uint(0), data.UserID)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestCancelOrderError", func(t *testing.T) {
		orderUseCase := NewOrderUseCase(mockOrderRepositoryError{}, mockCartRepositoryError{}, mockProductRepositoryError{})
		data, rows, err := orderUseCase.CancelOrder(_entities.OrdersDetail{}, 2, 2)
		assert.NotNil(t, err)
		assert.Nil(t, nil, data)
		assert.Nil(t, nil, rows)
	})
}






type mockOrderRepository struct{}

func (m mockOrderRepository) GetAllOrder(idToken int) ([]_entities.OrdersDetail, int, error) {
	return []_entities.OrdersDetail{
		{UserID: 1, ProductID: 1, TotalPrice: 5000, Status: "in process"},
	}, 1, nil
}

func (m mockOrderRepository) CreateOrder(request _entities.Order, orderCartID []uint) (_entities.Order, int, error) {
	return _entities.Order{
		UserID: 1, TotalPrice: 5000, StatusOrder: "not paid yet",
	}, 1, nil
}

func (m mockOrderRepository) GetHistoriOrderbyID(id int) (_entities.OrdersDetail, int, error) {
	return _entities.OrdersDetail{
		UserID: 1, ProductID: 1, TotalPrice: 5000, Status: "in process",
	}, 1, nil
}

func (m mockOrderRepository) CancelOrder(cancelOrder _entities.OrdersDetail) (_entities.OrdersDetail, int, error) {
	return _entities.OrdersDetail{
		UserID: 1, ProductID: 1, TotalPrice: 5000, Status: "in process",
	}, 1, nil
}

// === mock error ===

type mockOrderRepositoryError struct{}

func (m mockOrderRepositoryError) GetAllOrder(idToken int) ([]_entities.OrdersDetail, int, error) {
	return nil, 1, fmt.Errorf("error get all order")
}

func (m mockOrderRepositoryError) CreateOrder(request _entities.Order, orderCartID []uint) (_entities.Order, int, error) {
	return _entities.Order{}, 1, fmt.Errorf("error create order")
}

func (m mockOrderRepositoryError) GetHistoriOrderbyID(id int) (_entities.OrdersDetail, int, error) {
	return _entities.OrdersDetail{}, 0, fmt.Errorf("error get history")
}

func (m mockOrderRepositoryError) CancelOrder(cancelOrder _entities.OrdersDetail) (_entities.OrdersDetail, int, error) {
	return _entities.OrdersDetail{}, 0, fmt.Errorf("error cancel order")
}



// === mock success ===
type mockCartRepository struct{}

func (m mockCartRepository) GetAll(idToken int) ([]_entities.Cart, int, error) {
	return []_entities.Cart{
		{UserID: 1, ProductID: 1, Qty: 3, SubTotal: 5000},
	}, 1, nil
}

func (m mockCartRepository) GetCartById(id int) (_entities.Cart, int, error) {
	return _entities.Cart{
		UserID: 1, ProductID: 1, Qty: 3, SubTotal: 5000,
	}, 1, nil
}

func (m mockCartRepository) CreateCart(request _entities.Cart) (_entities.Cart, error) {
	return _entities.Cart{
		UserID: 1, ProductID: 1, Qty: 3, SubTotal: 5000,
	}, nil
}

func (m mockCartRepository) UpdateCart(request _entities.Cart) (_entities.Cart, int, error) {
	return _entities.Cart{
		UserID: 1, ProductID: 1, Qty: 3, SubTotal: 5000,
	}, 1, nil
}

func (m mockCartRepository) DeleteCart(id int) error {
	return nil
}


// === mock error ===

type mockCartRepositoryError struct{}

func (m mockCartRepositoryError) GetAll(idToken int) ([]_entities.Cart, int, error) {
	return nil, 1, fmt.Errorf("error get all cart")
}

func (m mockCartRepositoryError) GetCartById(id int) (_entities.Cart, int, error) {
	return _entities.Cart{}, 1, fmt.Errorf("error get all cart")
}

func (m mockCartRepositoryError) CreateCart(request _entities.Cart) (_entities.Cart, error) {
	return _entities.Cart{}, fmt.Errorf("error create cart")
}

func (m mockCartRepositoryError) UpdateCart(request _entities.Cart) (_entities.Cart, int, error) {
	return _entities.Cart{}, 0, fmt.Errorf("error update cart")
}

func (m mockCartRepositoryError) DeleteCart(id int) (error) {
	return fmt.Errorf("error delete cart")
}

// === mock success ===
type mockProductRepository struct{}

func (m mockProductRepository) GetProductById(id int) (_entities.Product, int, error) {
	return _entities.Product{
		ProductTitle: "product 1", ProductDesc: "desc", Price: 5000, UserID: 1, CatagoryID: 1, Stock: 2, UrlProduct: "url",
	}, 1, nil
}

func (m mockProductRepository) GetAll() ([]_entities.Product, error) {
	return []_entities.Product{
		{ProductTitle: "product 1", ProductDesc: "desc", Price: 5000, UserID: 1, CatagoryID: 1, Stock: 2, UrlProduct: "url"},
	}, nil
}

func (m mockProductRepository) GetAllProductUser(userID uint) ([]_entities.Product, error) {
	return []_entities.Product{
		{ProductTitle: "product 1", ProductDesc: "desc", Price: 5000, UserID: 1, CatagoryID: 1, Stock: 2, UrlProduct: "url"},
	}, nil
}

func (m mockProductRepository) CreateProduct(request _entities.Product) (_entities.Product, error) {
	return _entities.Product{
		ProductTitle: "product 1", ProductDesc: "desc", Price: 5000, UserID: 1, CatagoryID: 1, Stock: 2, UrlProduct: "url",
	}, nil
}

func (m mockProductRepository) UpdateProduct(request _entities.Product) (_entities.Product, int, error) {
	return _entities.Product{
		ProductTitle: "product 1", ProductDesc: "desc", Price: 5000, UserID: 1, CatagoryID: 1, Stock: 2, UrlProduct: "url",
	}, 1, nil
}

func (m mockProductRepository) DeleteProduct(id int) error {
	return nil
}


// === mock error ===

type mockProductRepositoryError struct{}

func (m mockProductRepositoryError) GetProductById(id int) (_entities.Product, int, error) {
	return _entities.Product{}, 0, fmt.Errorf("error get data user")
}

func (m mockProductRepositoryError) GetAll() ([]_entities.Product, error) {
	return nil, fmt.Errorf("error get all data user")
}

func (m mockProductRepositoryError) GetAllProductUser(userID uint) ([]_entities.Product, error) {
	return nil, fmt.Errorf("error get product user")
}

func (m mockProductRepositoryError) CreateProduct(request _entities.Product) (_entities.Product, error) {
	return _entities.Product{}, fmt.Errorf("error create user")
}

func (m mockProductRepositoryError) UpdateProduct(request _entities.Product) (_entities.Product, int, error) {
	return _entities.Product{}, 0, fmt.Errorf("error update data user")
}

func (m mockProductRepositoryError) DeleteProduct(id int) (error) {
	return fmt.Errorf("error update data user")
}


