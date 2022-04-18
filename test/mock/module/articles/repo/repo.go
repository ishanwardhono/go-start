// Code generated by MockGen. DO NOT EDIT.
// Source: ./module/articles/repo/repo.go

// Package mockarticlesrepo is a generated GoMock package.
package repo

import (
	entity "app/entity"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockArticleRepo is a mock of ArticleRepo interface.
type MockArticleRepo struct {
	ctrl     *gomock.Controller
	recorder *MockArticleRepoMockRecorder
}

// MockArticleRepoMockRecorder is the mock recorder for MockArticleRepo.
type MockArticleRepoMockRecorder struct {
	mock *MockArticleRepo
}

// NewMockArticleRepo creates a new mock instance.
func NewMockArticleRepo(ctrl *gomock.Controller) *MockArticleRepo {
	mock := &MockArticleRepo{ctrl: ctrl}
	mock.recorder = &MockArticleRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleRepo) EXPECT() *MockArticleRepoMockRecorder {
	return m.recorder
}

// GetAllArticle mocks base method.
func (m *MockArticleRepo) GetAllArticle(ctx context.Context) ([]entity.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllArticle", ctx)
	ret0, _ := ret[0].([]entity.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllArticle indicates an expected call of GetAllArticle.
func (mr *MockArticleRepoMockRecorder) GetAllArticle(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllArticle", reflect.TypeOf((*MockArticleRepo)(nil).GetAllArticle), ctx)
}

// GetArticle mocks base method.
func (m *MockArticleRepo) GetArticle(ctx context.Context, id int) (entity.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticle", ctx, id)
	ret0, _ := ret[0].(entity.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticle indicates an expected call of GetArticle.
func (mr *MockArticleRepoMockRecorder) GetArticle(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticle", reflect.TypeOf((*MockArticleRepo)(nil).GetArticle), ctx, id)
}

// InsertArticle mocks base method.
func (m *MockArticleRepo) InsertArticle(ctx context.Context, article entity.Article) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertArticle", ctx, article)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertArticle indicates an expected call of InsertArticle.
func (mr *MockArticleRepoMockRecorder) InsertArticle(ctx, article interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertArticle", reflect.TypeOf((*MockArticleRepo)(nil).InsertArticle), ctx, article)
}
