package usecase

import (
	"app/module/articles/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewArticlesFactory(t *testing.T) {
	factory := NewArticlesFactory(nil)
	assert.NotNil(t, factory)

	create := factory.Create(model.Article{})
	assert.NotNil(t, create)
	getAll := factory.GetAll()
	assert.NotNil(t, getAll)
	get := factory.Get(0)
	assert.NotNil(t, get)
}
