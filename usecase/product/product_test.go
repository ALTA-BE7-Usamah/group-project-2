package product

import (
	"fmt"
	_entities "group-project/limamart/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestGetProductById(t *testing.T) {
	t.Run("TestGetByIdSuccess", func(t *testing.T) {
		productUseCase := NewProductUseCase(mockProductRepository{}, mockCartRepository{})
		data, rows, err := productUseCase.GetProductById(1)
		assert.Nil(t, err)
		assert.Equal(t, "product 1", data.ProductTitle)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestGetProductByIdError", func(t *testing.T) {
		productUseCase := NewProductUseCase(mockProductRepositoryError{}, mockCartRepository{})
		data, rows, err := productUseCase.GetProductById(1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, rows)
		assert.Equal(t, _entities.Product{}, data)
	})
}

func TestCreateProduct(t *testing.T) {
	t.Run("TestCreateProductSuccess", func(t *testing.T) {
		productUseCase := NewProductUseCase(mockProductRepository{}, mockCartRepository{})
		data, err := productUseCase.CreateProduct(_entities.Product{CatagoryID: 1})
		assert.Nil(t, nil, err)
		assert.Equal(t, uint(1), data.CatagoryID)
	})

	t.Run("TestCreateProductSuccess", func(t *testing.T) {
		productUseCase := NewProductUseCase(mockProductRepository{}, mockCartRepository{})
		data, err := productUseCase.CreateProduct(_entities.Product{CatagoryID: 1, ProductTitle: "product 1"})
		assert.Nil(t, nil, err)
		assert.Equal(t, uint(1), data.CatagoryID)
	})

	t.Run("TestCreateProductSuccess", func(t *testing.T) {
		productUseCase := NewProductUseCase(mockProductRepository{}, mockCartRepository{})
		data, err := productUseCase.CreateProduct(_entities.Product{CatagoryID: 1, ProductTitle: "product 1", ProductDesc: "desc"})
		assert.Nil(t, nil, err)
		assert.Equal(t, uint(1), data.CatagoryID)
	})

	t.Run("TestCreateProductSuccess", func(t *testing.T) {
		productUseCase := NewProductUseCase(mockProductRepository{}, mockCartRepository{})
		data, err := productUseCase.CreateProduct(_entities.Product{CatagoryID: 1, ProductTitle: "product 1", ProductDesc: "desc", Price: 5000})
		assert.Nil(t, nil, err)
		assert.Equal(t, uint(1), data.CatagoryID)
	})

	t.Run("TestCreateProductSuccess", func(t *testing.T) {
		productUseCase := NewProductUseCase(mockProductRepository{}, mockCartRepository{})
		data, err := productUseCase.CreateProduct(_entities.Product{CatagoryID: 1, ProductTitle: "product 1", ProductDesc: "desc", Price: 5000, Stock: 3})
		assert.Nil(t, nil, err)
		assert.Equal(t, uint(1), data.CatagoryID)
	})

	t.Run("TestCreateProductSuccess", func(t *testing.T) {
		productUseCase := NewProductUseCase(mockProductRepository{}, mockCartRepository{})
		data, err := productUseCase.CreateProduct(_entities.Product{CatagoryID: 1, ProductTitle: "product 1", ProductDesc: "desc", Price: 5000, Stock: 3, UrlProduct: "url"})
		assert.Nil(t, nil, err)
		assert.Equal(t, uint(1), data.CatagoryID)
	})


	t.Run("TestGetProductByIdError", func(t *testing.T) {
		productUseCase := NewProductUseCase(mockProductRepositoryError{}, mockCartRepository{})
		data, err := productUseCase.CreateProduct(_entities.Product{})
		assert.NotNil(t, err)
		assert.Equal(t, _entities.Product{}, data)
		assert.Nil(t, nil, err)
	})
}

func TestUpdateProduct(t *testing.T) {
	t.Run("TestUpdateProductSuccess", func(t *testing.T) {
		productUseCase := NewProductUseCase(mockProductRepository{}, mockCartRepository{})
		data, rows, err := productUseCase.UpdateProduct(_entities.Product{CatagoryID: 1}, 1, 1)
		assert.Nil(t, err)
		assert.Equal(t, uint(1), data.CatagoryID)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestUpdateProductSuccess", func(t *testing.T) {
		productUseCase := NewProductUseCase(mockProductRepository{}, mockCartRepository{})
		data, rows, err := productUseCase.UpdateProduct(_entities.Product{CatagoryID: 1, ProductTitle: "product 1"}, 1, 1)
		assert.Nil(t, err)
		assert.Equal(t, uint(1), data.CatagoryID)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestUpdateProductSuccess", func(t *testing.T) {
		productUseCase := NewProductUseCase(mockProductRepository{}, mockCartRepository{})
		data, rows, err := productUseCase.UpdateProduct(_entities.Product{CatagoryID: 1, ProductTitle: "product 1", ProductDesc: "desc"}, 1, 1)
		assert.Nil(t, err)
		assert.Equal(t, uint(1), data.CatagoryID)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestUpdateProductSuccess", func(t *testing.T) {
		productUseCase := NewProductUseCase(mockProductRepository{}, mockCartRepository{})
		data, rows, err := productUseCase.UpdateProduct(_entities.Product{CatagoryID: 1, ProductTitle: "product 1", ProductDesc: "desc", Price: 5000}, 1, 1)
		assert.Nil(t, err)
		assert.Equal(t, uint(1), data.CatagoryID)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestUpdateProductSuccess", func(t *testing.T) {
		productUseCase := NewProductUseCase(mockProductRepository{}, mockCartRepository{})
		data, rows, err := productUseCase.UpdateProduct(_entities.Product{CatagoryID: 1, ProductTitle: "product 1", ProductDesc: "desc", Price: 5000, Stock: 5}, 1, 1)
		assert.Nil(t, err)
		assert.Equal(t, uint(1), data.CatagoryID)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestUpdateProductSuccess", func(t *testing.T) {
		productUseCase := NewProductUseCase(mockProductRepository{}, mockCartRepository{})
		data, rows, err := productUseCase.UpdateProduct(_entities.Product{CatagoryID: 1, ProductTitle: "product 1", ProductDesc: "desc", Price: 5000, Stock: 5, UrlProduct: "url"}, 1, 1)
		assert.Nil(t, err)
		assert.Equal(t, uint(1), data.CatagoryID)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestUpdateProductError", func(t *testing.T) {
		productUseCase := NewProductUseCase(mockProductRepositoryError{}, mockCartRepository{})
		data, rows, err := productUseCase.UpdateProduct(_entities.Product{}, 1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, rows)
		assert.Equal(t, _entities.Product{}, data)
	})
}

func TestDeleteProduct(t *testing.T) {
	t.Run("TestDeleteProductSuccess", func(t *testing.T) {
		productUseCase := NewProductUseCase(mockProductRepository{}, mockCartRepository{})
		err := productUseCase.DeleteProduct(1)
		assert.Nil(t, err)
	})

	t.Run("TestGetProductByIdError", func(t *testing.T) {
		productUseCase := NewProductUseCase(mockProductRepositoryError{}, mockCartRepository{})
		err := productUseCase.DeleteProduct(1)
		assert.NotNil(t, err)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		productUseCase := NewProductUseCase(mockProductRepository{}, mockCartRepository{})
		data, err := productUseCase.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, "product 1", data[0].ProductTitle)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		productUseCase := NewProductUseCase(mockProductRepositoryError{}, mockCartRepository{})
		data, err := productUseCase.GetAll()
		assert.NotNil(t, err)
		assert.Nil(t, data)
	})
}

func TestGetAllProductUser(t *testing.T) {
	t.Run("TestGetAllProductUserSuccess", func(t *testing.T) {
		productUseCase := NewProductUseCase(mockProductRepository{}, mockCartRepository{})
		data, err := productUseCase.GetAllProductUser(1)
		assert.Nil(t, err)
		assert.Equal(t, "product 1", data[0].ProductTitle)
	})

	t.Run("TestGetAllProductUserError", func(t *testing.T) {
		productUseCase := NewProductUseCase(mockProductRepositoryError{}, mockCartRepository{})
		data, err := productUseCase.GetAllProductUser(1)
		assert.NotNil(t, err)
		assert.Nil(t, data)
	})
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

func (m mockProductRepository) DeleteProduct(id int, cart []_entities.Cart) error {
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

func (m mockProductRepositoryError) DeleteProduct(id int, cart []_entities.Cart) (error) {
	return fmt.Errorf("error update data user")
}

// === mock success ===
type mockCartRepository struct{}

func (m mockCartRepository) GetCartByProductId(idProduct int) ([]_entities.Cart, int, error) {
	return []_entities.Cart{
		{UserID: 1, ProductID: 1, Qty: 3, SubTotal: 5000},
	}, 1, nil
}
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

func (m mockCartRepositoryError) GetCartByProductId(idProduct int) ([]_entities.Cart, int, error) {
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

