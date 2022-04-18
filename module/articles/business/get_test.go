package usecase

import (
	"app/test/mock/module/articles/repo"
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_getArticle_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockArticleRepo := repo.NewMockArticleRepo(ctrl)
	mockData := repo.GetMockData()

	tests := []struct {
		name    string
		req     int
		mock    func()
		want    interface{}
		wantErr bool
	}{
		{
			name: "success",
			req:  1,
			mock: func() {
				mockArticleRepo.EXPECT().GetArticle(gomock.Any(), gomock.Any()).Return(mockData[0], nil)
			},
			want: mockData[0],
		},
		{
			name:    "failed, validation",
			req:     0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}
			m := &getArticle{
				Req:         tt.req,
				RepoArticle: mockArticleRepo,
			}
			got, err := m.Execute(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("getArticle.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getArticle.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
