package usecase

import (
	"app/entity"
	"app/mock/database/repo/users"
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_createUser_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepo := users.NewMockUserRepo(ctrl)

	tests := []struct {
		name    string
		req     entity.User
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
			req: entity.User{
				Name:  "ishan",
				Email: "ishan@mail . com",
			},
			mock: func() {
				mockUserRepo.EXPECT().InsertUser(gomock.Any()).Return(errors.New("error execute repo"))
			},
			want:    "insert failed",
			wantErr: true,
		},
		{
			name: "success",
			req: entity.User{
				Name:  "ishan",
				Email: "ishan@mail . com",
			},
			mock: func() {
				mockUserRepo.EXPECT().InsertUser(gomock.Any()).Return(nil)
			},
			want: "insert success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}
			m := &createUser{
				Req:      tt.req,
				RepoUser: mockUserRepo,
			}
			got, err := m.Execute(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("createUser.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createUser.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
