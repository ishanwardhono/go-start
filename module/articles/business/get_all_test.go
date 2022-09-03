package usecase

import (
	"app/module/articles/model"
	"app/test/mock/module/articles/repo"
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_getArticles_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockArticleRepo := repo.NewMockArticleRepo(ctrl)
	mockData := repo.GetMockData()

	tests := []struct {
		name    string
		req     model.Article
		mock    func()
		want    interface{}
		wantErr bool
	}{
		{
			name: "success",
			req: model.Article{
				Title:   "test title",
				Content: "test article content",
				Author:  "test author",
			},
			mock: func() {
				mockArticleRepo.EXPECT().GetAllArticle(gomock.Any()).Return(mockData, nil)
			},
			want: mockData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}
			m := &getArticles{
				RepoArticle: mockArticleRepo,
			}
			got, err := m.Execute(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("getArticles.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getArticles.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
