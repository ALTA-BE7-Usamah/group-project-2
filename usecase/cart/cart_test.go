package cart

import (
	"fmt"
	_entities "group-project/limamart/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	t.Run("TestCreateCartSuccess", func(t *testing.T) {
		cartUseCase := NewCartUseCase(mockCartRepository{}, mockProductRepository{})
		data, rows, err := cartUseCase.GetAll(1)
		assert.Nil(t, nil, err)
		assert.Equal(t, uint(1), data[0].UserID)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestCreateCartError", func(t *testing.T) {
		cartUseCase := NewCartUseCase(mockCartRepositoryError{}, mockProductRepositoryError{})
		data, rows, err := cartUseCase.GetAll(1)
		assert.NotNil(t, err)
		assert.Nil(t, data)
		assert.Nil(t, nil, rows)
	})
}

func TestCreateCart(t *testing.T) {
	t.Run("TestCreateCartSuccess", func(t *testing.T) {
		cartUseCase := NewCartUseCase(mockCartRepository{}, mockProductRepository{})
		data, err := cartUseCase.CreateCart(_entities.Cart{})
		assert.Nil(t, nil, err)
		assert.Equal(t, uint(0), data.ProductID)
	})

	t.Run("TestCreateCartError", func(t *testing.T) {
		cartUseCase := NewCartUseCase(mockCartRepositoryError{}, mockProductRepositoryError{})
		data, err := cartUseCase.CreateCart(_entities.Cart{})
		assert.NotNil(t, err)
		assert.Equal(t, _entities.Cart{}, data)
	})
}

func TestUpdateCart(t *testing.T) {
	t.Run("TestCreateCartSuccess", func(t *testing.T) {
		cartUseCase := NewCartUseCase(mockCartRepository{}, mockProductRepository{})
		data, rows, err := cartUseCase.UpdateCart(1, 1, _entities.Cart{})
		assert.Nil(t, nil, err)
		assert.Equal(t, uint(1), data.UserID)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestCreateCartError", func(t *testing.T) {
		cartUseCase := NewCartUseCase(mockCartRepositoryError{}, mockProductRepositoryError{})
		data, rows, err := cartUseCase.UpdateCart(1, 2, _entities.Cart{})
		assert.NotNil(t, err)
		assert.Equal(t, 0, rows)
		assert.Equal(t, _entities.Cart{}, data)
	})
}

func TestDeeleteCart(t *testing.T) {
	t.Run("TestDeleteCartSuccess", func(t *testing.T) {
		cartUseCase := NewCartUseCase(mockCartRepository{}, mockProductRepository{})
		err := cartUseCase.DeleteCart(1)
		assert.Nil(t, nil, err)
	})

	t.Run("TestDeleteCartError", func(t *testing.T) {
		cartUseCase := NewCartUseCase(mockCartRepositoryError{}, mockProductRepositoryError{})
		err := cartUseCase.DeleteCart(1)
		assert.NotNil(t, err)
	})
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

