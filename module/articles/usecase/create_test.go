package usecase

import (
	corehttp "app/core/handler/http"
	"app/entity"
	"app/test/mock/module/articles/repo"
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_createArticle_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockArticleRepo := repo.NewMockArticleRepo(ctrl)

	tests := []struct {
		name    string
		req     entity.Article
		mock    func()
		want    interface{}
		wantErr bool
	}{
		{
			name:    "error validation",
			wantErr: true,
		},
		{
			name: "failed execute repo",
			req: entity.Article{
				Title:   "test title",
				Content: "test article content",
				Author:  "test author",
			},
			mock: func() {
				mockArticleRepo.EXPECT().InsertArticle(gomock.Any(), gomock.Any()).Return(0, errors.New("error execute repo"))
			},
			want:    "insert failed",
			wantErr: true,
		},
		{
			name: "success",
			req: entity.Article{
				Title:   "test title",
				Content: "test article content",
				Author:  "test author",
			},
			mock: func() {
				mockArticleRepo.EXPECT().InsertArticle(gomock.Any(), gomock.Any()).Return(1, nil)
			},
			want: corehttp.Response{
				StatusCode: http.StatusCreated,
				Message:    "Success",
				Data: map[string]interface{}{
					"id": 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}
			m := &createArticle{
				Req:         tt.req,
				RepoArticle: mockArticleRepo,
			}
			got, err := m.Execute(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("createArticle.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createArticle.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
