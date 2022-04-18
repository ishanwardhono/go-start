package usecase

import (
	"app/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewArticlesFactory(t *testing.T) {
	factory := NewArticlesFactory(nil)
	assert.NotNil(t, factory)

	create := factory.Create(entity.Article{})
	assert.NotNil(t, create)
	getAll := factory.GetAll()
	assert.NotNil(t, getAll)
	get := factory.Get(0)
	assert.NotNil(t, get)
}
