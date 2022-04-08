package product

import (
	"fmt"
	_entities "group-project/limamart/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestGetAllCatagory(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		categoryUseCase := NewCatagoryUseCase(mockCategoryRepository{})
		data, err := categoryUseCase.GetAllCatagory()
		assert.Nil(t, err)
		assert.Equal(t, "category 1", data[0].CatagoryName)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		categoryUseCase := NewCatagoryUseCase(mockCategoryRepositoryError{})
		data, err := categoryUseCase.GetAllCatagory()
		assert.NotNil(t, err)
		assert.Nil(t, data)
	})
}


// === mock success ===
type mockCategoryRepository struct{}


func (m mockCategoryRepository) GetAllCatagory() ([]_entities.Catagory, error) {
	return []_entities.Catagory{
		{CatagoryName: "category 1"},
	}, nil
}


// === mock error ===

type mockCategoryRepositoryError struct{}

func (m mockCategoryRepositoryError) GetAllCatagory() ([]_entities.Catagory, error) {
	return nil, fmt.Errorf("error get all data user")
}

